package portrait

import (
	"image"
	"image/color"
	"image/draw"
	"io"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type balloonParams struct {
	border    int
	cornerCut int
	bodyW     int
	bodyH     int
	tailW     int
	tailH     int
	tailX     int
	totalH    int
}

func newBalloonParams(innerW, innerH, border int) balloonParams {
	cornerCut := border * 2
	bodyW := innerW + border*2
	bodyH := innerH + border*2
	tailW := border * 3
	tailH := border * 4
	return balloonParams{
		border:    border,
		cornerCut: cornerCut,
		bodyW:     bodyW,
		bodyH:     bodyH,
		tailW:     tailW,
		tailH:     tailH,
		tailX:     cornerCut,
		totalH:    bodyH + tailH,
	}
}

func allocGrid(rows, cols int) [][]int {
	grid := make([][]int, rows)
	for y := range grid {
		grid[y] = make([]int, cols)
	}
	return grid
}

func isBodyCorner(x, y int, p balloonParams) bool {
	return (x < p.cornerCut && y < p.cornerCut) ||
		(x >= p.bodyW-p.cornerCut && y < p.cornerCut) ||
		(x < p.cornerCut && y >= p.bodyH-p.cornerCut) ||
		(x >= p.bodyW-p.cornerCut && y >= p.bodyH-p.cornerCut)
}

func isBodyEdge(x, y int, p balloonParams) bool {
	return x < p.border || x >= p.bodyW-p.border || y < p.border || y >= p.bodyH-p.border
}

func bodyPixelValue(x, y int, p balloonParams) int {
	if isBodyCorner(x, y, p) {
		return 0
	}
	if isBodyEdge(x, y, p) {
		return 1
	}
	return 2
}

func fillBodyGrid(grid [][]int, p balloonParams) {
	for y := 0; y < p.bodyH; y++ {
		for x := 0; x < p.bodyW; x++ {
			grid[y][x] = bodyPixelValue(x, y, p)
		}
	}
}

func openTailGap(grid [][]int, p balloonParams) {
	for x := p.tailX + p.border; x < p.tailX+p.tailW-p.border && x < p.bodyW-p.cornerCut; x++ {
		for y := p.bodyH - p.border; y < p.bodyH; y++ {
			if grid[y][x] == 1 {
				grid[y][x] = 2
			}
		}
	}
}

func isTailEdge(x, y int, p balloonParams) bool {
	return x < p.tailX+p.border || x >= p.tailX+p.tailW-p.border || y >= p.totalH-p.border
}

func fillTailGrid(grid [][]int, p balloonParams) {
	for y := p.bodyH; y < p.totalH; y++ {
		for x := p.tailX; x < p.tailX+p.tailW && x < p.bodyW; x++ {
			if isTailEdge(x, y, p) {
				grid[y][x] = 1
			} else {
				grid[y][x] = 2
			}
		}
	}
}

func generateBalloonGrid(innerW, innerH, border int) [][]int {
	p := newBalloonParams(innerW, innerH, border)
	grid := allocGrid(p.totalH, p.bodyW)
	fillBodyGrid(grid, p)
	openTailGap(grid, p)
	fillTailGrid(grid, p)
	return grid
}

func gridSize(grid [][]int) (int, int) {
	h := len(grid)
	if h == 0 {
		return 0, 0
	}
	return len(grid[0]), h
}

func balloonBorderColor(opt Options) color.Color {
	if opt.BalloonBorderColor != nil {
		return opt.BalloonBorderColor
	}
	return color.RGBA{R: 0, G: 0, B: 0, A: 255}
}

func balloonFillColor(opt Options) color.Color {
	if opt.BalloonFillColor != nil {
		return opt.BalloonFillColor
	}
	return color.RGBA{R: 255, G: 255, B: 255, A: 255}
}

func balloonTextColor(opt Options) color.Color {
	if opt.TextColor != nil && opt.TextColor.A != 0 {
		return opt.TextColor
	}
	return color.RGBA{R: 0, G: 0, B: 0, A: 255}
}

func balloonBorderThickness(multiple int) int {
	if multiple < 1 {
		return 1
	}
	return multiple
}

func balloonPadding(border int) int {
	return border * 4
}

func measureBalloonText(face font.Face, text string) (int, int) {
	metrics := face.Metrics()
	textWidth := font.MeasureString(face, text).Ceil()
	textHeight := (metrics.Ascent + metrics.Descent).Ceil()
	return textWidth, textHeight
}

func balloonInnerSize(textWidth, textHeight, padding int) (int, int) {
	return textWidth + padding*2, textHeight + padding*2
}

func newBalloonCanvas(portraitSize, gap, gridW, gridH int, bgColor color.Color, portrait *image.Paletted) *image.NRGBA {
	canvasH := portraitSize
	if gridH > canvasH {
		canvasH = gridH
	}
	canvas := image.NewNRGBA(image.Rect(0, 0, portraitSize+gap+gridW, canvasH))
	draw.Draw(canvas, canvas.Bounds(), image.NewUniform(bgColor), image.Point{}, draw.Src)
	draw.Draw(canvas, portrait.Bounds(), portrait, image.Point{}, draw.Over)
	return canvas
}

func balloonColorPalette(border, fill color.Color) []color.Color {
	return []color.Color{
		color.RGBA{A: 0},
		border,
		fill,
	}
}

func renderBalloonGrid(canvas *image.NRGBA, grid [][]int, bx, by int, palette []color.Color) {
	for y, row := range grid {
		for x, v := range row {
			if v != 0 {
				canvas.Set(bx+x, by+y, palette[v])
			}
		}
	}
}

func drawBalloonText(canvas *image.NRGBA, face font.Face, text string, bx, by, border, padding int, textColor color.Color) {
	metrics := face.Metrics()
	textX := bx + border + padding
	textY := by + border + padding + metrics.Ascent.Ceil()
	d := &font.Drawer{
		Dst:  canvas,
		Src:  image.NewUniform(textColor),
		Face: face,
		Dot:  fixed.P(textX, textY),
	}
	d.DrawString(text)
}

func (p *Portrait) drawBalloon(portrait *image.Paletted) image.Image {
	face := p.newFontFace()
	if closer, ok := face.(io.Closer); ok {
		defer closer.Close()
	}

	border := balloonBorderThickness(p.opt.Multiple)
	padding := balloonPadding(border)
	textWidth, textHeight := measureBalloonText(face, p.opt.Text)
	innerW, innerH := balloonInnerSize(textWidth, textHeight, padding)

	grid := generateBalloonGrid(innerW, innerH, border)
	gridW, gridH := gridSize(grid)

	gap := padding
	canvas := newBalloonCanvas(p.opt.Size, gap, gridW, gridH, p.opt.BackgroundColor, portrait)

	palette := balloonColorPalette(balloonBorderColor(p.opt), balloonFillColor(p.opt))
	bx := p.opt.Size + gap
	renderBalloonGrid(canvas, grid, bx, 0, palette)
	drawBalloonText(canvas, face, p.opt.Text, bx, 0, border, padding, balloonTextColor(p.opt))

	return canvas
}
