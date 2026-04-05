package portrait

import (
	"image"
	"image/color"
	"image/draw"
	"io"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

// generateBalloonGrid creates a 2D pixel grid for an 8-bit-style speech balloon.
// Grid values: 0=transparent, 1=border/line color, 2=fill/background color.
// The body is a rounded rectangle with 8-bit corner cuts.
// A rectangular tail extends downward from the bottom of the body.
//
//	innerW: inner content area width (pixels)
//	innerH: inner content area height (pixels)
//	border: border thickness (pixels)
func generateBalloonGrid(innerW, innerH, border int) [][]int {
	cornerCut := border * 2
	bodyW := innerW + border*2
	bodyH := innerH + border*2
	tailW := border * 3
	tailH := border * 4
	tailX := cornerCut
	totalH := bodyH + tailH

	grid := make([][]int, totalH)
	for y := range grid {
		grid[y] = make([]int, bodyW)
	}

	// Draw body with 8-bit corner cuts
	for y := 0; y < bodyH; y++ {
		for x := 0; x < bodyW; x++ {
			inCorner := (x < cornerCut && y < cornerCut) ||
				(x >= bodyW-cornerCut && y < cornerCut) ||
				(x < cornerCut && y >= bodyH-cornerCut) ||
				(x >= bodyW-cornerCut && y >= bodyH-cornerCut)
			if inCorner {
				grid[y][x] = 0
			} else if x < border || x >= bodyW-border || y < border || y >= bodyH-border {
				grid[y][x] = 1
			} else {
				grid[y][x] = 2
			}
		}
	}

	// Open the bottom border where the tail connects (inner area only)
	for x := tailX + border; x < tailX+tailW-border && x < bodyW-cornerCut; x++ {
		for y := bodyH - border; y < bodyH; y++ {
			if grid[y][x] == 1 {
				grid[y][x] = 2
			}
		}
	}

	// Draw tail
	for y := bodyH; y < totalH; y++ {
		for x := tailX; x < tailX+tailW && x < bodyW; x++ {
			isLeftEdge := x < tailX+border
			isRightEdge := x >= tailX+tailW-border
			isBottomEdge := y >= totalH-border
			if isLeftEdge || isRightEdge || isBottomEdge {
				grid[y][x] = 1
			} else {
				grid[y][x] = 2
			}
		}
	}

	return grid
}

// drawBalloon renders a speech balloon containing the text, placed to the right of the portrait.
// Default balloon colors: black border, white fill.
func (p *Portrait) drawBalloon(portrait *image.Paletted) image.Image {
	face := p.newFontFace()
	if closer, ok := face.(io.Closer); ok {
		defer closer.Close()
	}

	// Balloon colors: default black border, white fill
	var borderColor color.Color = color.RGBA{R: 0, G: 0, B: 0, A: 255}
	var fillColor color.Color = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	if p.opt.BalloonBorderColor != nil {
		borderColor = p.opt.BalloonBorderColor
	}
	if p.opt.BalloonFillColor != nil {
		fillColor = p.opt.BalloonFillColor
	}

	// Text color defaults to black (readable on white balloon).
	// A nil or fully-transparent TextColor means the user did not set a color,
	// so we fall back to black. If the user explicitly sets a color, it is used.
	var textColor color.Color = color.RGBA{R: 0, G: 0, B: 0, A: 255}
	if p.opt.TextColor != nil && p.opt.TextColor.A != 0 {
		textColor = p.opt.TextColor
	}

	metrics := face.Metrics()
	textWidth := font.MeasureString(face, p.opt.Text).Ceil()
	textHeight := (metrics.Ascent + metrics.Descent).Ceil()

	border := p.opt.Multiple
	if border < 1 {
		border = 1
	}
	padding := border * 4

	innerW := textWidth + padding*2
	innerH := textHeight + padding*2

	grid := generateBalloonGrid(innerW, innerH, border)
	gridH := len(grid)
	gridW := 0
	if gridH > 0 {
		gridW = len(grid[0])
	}

	gap := padding
	canvasW := p.opt.Size + gap + gridW
	canvasH := p.opt.Size
	if gridH > canvasH {
		canvasH = gridH
	}

	canvas := image.NewNRGBA(image.Rect(0, 0, canvasW, canvasH))
	draw.Draw(canvas, canvas.Bounds(), image.NewUniform(p.opt.BackgroundColor), image.Point{}, draw.Src)
	draw.Draw(canvas, portrait.Bounds(), portrait, image.Point{}, draw.Over)

	// Render balloon grid
	palette := []color.Color{
		color.RGBA{0, 0, 0, 0}, // 0: transparent
		borderColor,             // 1: border
		fillColor,               // 2: fill
	}

	bx := p.opt.Size + gap
	by := 0
	for y, row := range grid {
		for x, v := range row {
			if v != 0 {
				canvas.Set(bx+x, by+y, palette[v])
			}
		}
	}

	// Draw text inside balloon body
	textX := bx + border + padding
	textY := by + border + padding + metrics.Ascent.Ceil()

	d := &font.Drawer{
		Dst:  canvas,
		Src:  image.NewUniform(textColor),
		Face: face,
		Dot:  fixed.P(textX, textY),
	}
	d.DrawString(p.opt.Text)

	return canvas
}
