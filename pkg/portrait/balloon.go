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
	bodyX     int
	tailW     int
	tailH     int
	tailY     int
	totalW    int
}

func newBalloonParams(innerW, innerH, border int) balloonParams {
	cornerCut := border * 2
	bodyW := innerW + border*2
	bodyH := innerH + border*2
	tailW := border * 2
	tailH := tailW*2 - 1
	tailY := (bodyH - tailH) / 2
	return balloonParams{
		border:    border,
		cornerCut: cornerCut,
		bodyW:     bodyW,
		bodyH:     bodyH,
		bodyX:     tailW,
		tailW:     tailW,
		tailH:     tailH,
		tailY:     tailY,
		totalW:    tailW + bodyW,
	}
}

func allocGrid(rows, cols int) [][]int {
	grid := make([][]int, rows)
	for y := range grid {
		grid[y] = make([]int, cols)
	}
	return grid
}

func isBodyCorner(bx, y int, p balloonParams) bool {
	return (bx < p.cornerCut && y < p.cornerCut) ||
		(bx >= p.bodyW-p.cornerCut && y < p.cornerCut) ||
		(bx < p.cornerCut && y >= p.bodyH-p.cornerCut) ||
		(bx >= p.bodyW-p.cornerCut && y >= p.bodyH-p.cornerCut)
}

func isBodyEdge(bx, y int, p balloonParams) bool {
	return bx < p.border || bx >= p.bodyW-p.border || y < p.border || y >= p.bodyH-p.border
}

func bodyPixelValue(bx, y int, p balloonParams) int {
	if isBodyCorner(bx, y, p) {
		return 0
	}
	if isBodyEdge(bx, y, p) {
		return 1
	}
	return 2
}

func fillBodyGrid(grid [][]int, p balloonParams) {
	for y := 0; y < p.bodyH; y++ {
		for bx := 0; bx < p.bodyW; bx++ {
			grid[y][p.bodyX+bx] = bodyPixelValue(bx, y, p)
		}
	}
}

func openTailGap(grid [][]int, p balloonParams) {
	for y := p.tailY + p.border; y < p.tailY+p.tailH-p.border; y++ {
		for bx := 0; bx < p.border; bx++ {
			if grid[y][p.bodyX+bx] == 1 {
				grid[y][p.bodyX+bx] = 2
			}
		}
	}
}

func tailMidY(p balloonParams) int {
	return p.tailY + p.tailH/2
}

func tailHalfSpan(x int) int {
	return x
}

func isTailPixel(x, y int, p balloonParams) bool {
	mid := tailMidY(p)
	half := tailHalfSpan(x)
	return y >= mid-half && y <= mid+half
}

func isTailBorderPixel(x, y int, p balloonParams) bool {
	mid := tailMidY(p)
	half := tailHalfSpan(x)
	return y < mid-half+p.border || y > mid+half-p.border
}

func tailPixelValue(x, y int, p balloonParams) int {
	if !isTailPixel(x, y, p) {
		return 0
	}
	if isTailBorderPixel(x, y, p) {
		return 1
	}
	return 2
}

func fillTailGrid(grid [][]int, p balloonParams) {
	for y := p.tailY; y < p.tailY+p.tailH; y++ {
		for x := 0; x < p.tailW; x++ {
			v := tailPixelValue(x, y, p)
			if v != 0 {
				grid[y][x] = v
			}
		}
	}
}

func generateBalloonGrid(innerW, innerH, border int) [][]int {
	p := newBalloonParams(innerW, innerH, border)
	grid := allocGrid(p.bodyH, p.totalW)
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

func newBalloonCanvas(portraitSize, balloonW, balloonH int, bgColor color.Color, portrait *image.Paletted) *image.NRGBA {
	canvasH := portraitSize
	if balloonH > canvasH {
		canvasH = balloonH
	}
	canvas := image.NewNRGBA(image.Rect(0, 0, portraitSize+balloonW, canvasH))
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

func drawBalloonText(canvas *image.NRGBA, face font.Face, text string, bx, by, bodyX, border, padding int, textColor color.Color) {
	metrics := face.Metrics()
	textX := bx + bodyX + border + padding
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

	params := newBalloonParams(innerW, innerH, border)
	grid := generateBalloonGrid(innerW, innerH, border)
	balloonW, balloonH := gridSize(grid)

	canvas := newBalloonCanvas(p.opt.Size, balloonW, balloonH, p.opt.BackgroundColor, portrait)

	palette := balloonColorPalette(balloonBorderColor(p.opt), balloonFillColor(p.opt))
	balloonX := p.opt.Size
	renderBalloonGrid(canvas, grid, balloonX, 0, palette)
	drawBalloonText(canvas, face, p.opt.Text, balloonX, 0, params.bodyX, border, padding, balloonTextColor(p.opt))

	return canvas
}
