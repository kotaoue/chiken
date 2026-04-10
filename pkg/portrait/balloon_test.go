package portrait

import (
	"image/color"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateBalloonGrid(t *testing.T) {
	border := 1
	innerW := 20
	innerH := 10
	grid := generateBalloonGrid(innerW, innerH, border)

	bodyW := innerW + border*2
	bodyH := innerH + border*2
	tailW := border * 4
	tailH := border * 3
	totalW := tailW + bodyW

	assert.Equal(t, bodyH, len(grid), "grid height should equal bodyH")
	for _, row := range grid {
		assert.Equal(t, totalW, len(row), "each grid row width should match tailW+bodyW")
	}

	// Body corners (body-local coords) should be transparent
	cornerCut := border * 2
	bodyX := tailW
	assert.Equal(t, 0, grid[0][bodyX], "body top-left corner should be transparent")
	assert.Equal(t, 0, grid[0][bodyX+bodyW-1], "body top-right corner should be transparent")
	assert.Equal(t, 0, grid[bodyH-1][bodyX], "body bottom-left corner should be transparent")
	assert.Equal(t, 0, grid[bodyH-1][bodyX+bodyW-1], "body bottom-right corner should be transparent")

	// Inside corner cut area should be transparent
	if cornerCut > 1 {
		assert.Equal(t, 0, grid[0][bodyX+1], "body top-left inner corner should be transparent")
	}

	// Body border cells (non-corner) should be 1
	assert.Equal(t, 1, grid[0][bodyX+cornerCut], "body top border should be 1")
	assert.Equal(t, 1, grid[cornerCut][bodyX+bodyW-1], "body right border should be 1")

	// Inner body cells should be fill (2)
	assert.Equal(t, 2, grid[cornerCut][bodyX+cornerCut], "inner body cell should be fill (2)")

	// Tail cells should be present (1 or 2) in tail area
	tailY := (bodyH - tailH) / 2
	tailFound := false
	for y := tailY; y < tailY+tailH; y++ {
		for x := 0; x < tailW; x++ {
			if grid[y][x] != 0 {
				tailFound = true
			}
		}
	}
	assert.True(t, tailFound, "tail cells should be present")

	// Tail left tip should be border color
	assert.Equal(t, 1, grid[tailY][0], "tail left tip top row should be border (1)")

	// Outside tail area should be transparent
	assert.Equal(t, 0, grid[0][0], "above tail area should be transparent")
}

func TestGenerateBalloonGrid_Border2(t *testing.T) {
	border := 2
	innerW := 40
	innerH := 20
	grid := generateBalloonGrid(innerW, innerH, border)

	bodyW := innerW + border*2
	bodyH := innerH + border*2
	tailW := border * 4
	totalW := tailW + bodyW

	assert.Equal(t, bodyH, len(grid), "grid height should equal bodyH")
	for _, row := range grid {
		assert.Equal(t, totalW, len(row), "each grid row width should match tailW+bodyW")
	}

	// Body corners should be transparent
	cornerCut := border * 2
	bodyX := tailW
	assert.Equal(t, 0, grid[0][bodyX], "body top-left corner should be transparent")
	assert.Equal(t, 0, grid[0][bodyX+cornerCut-1], "body top-left inner corner should be transparent")

	// Non-corner top body border should be 1
	assert.Equal(t, 1, grid[0][bodyX+cornerCut], "body top border should be 1")

	// Inner body cells should be fill (2)
	assert.Equal(t, 2, grid[cornerCut][bodyX+cornerCut], "inner body cell should be fill (2)")
}

func TestPortrait_DrawBalloon(t *testing.T) {
	// Test balloon encoding via Encode()
	opts := Options{
		Size:               32,
		BaseSize:           32,
		Multiple:           1,
		Style:              "basic",
		Theme:              "white",
		BackgroundColor:    &color.RGBA{R: 0, G: 0, B: 0, A: 255},
		Format:             "png",
		Output:             io.Discard,
		Text:               "Hello!",
		Balloon:            true,
		BalloonBorderColor: &color.RGBA{R: 0, G: 0, B: 0, A: 255},
		BalloonFillColor:   &color.RGBA{R: 255, G: 255, B: 255, A: 255},
	}
	p := NewPortrait(opts)
	err := p.Encode()
	assert.NoError(t, err, "Portrait.Encode() with balloon should not fail")
}

func TestPortrait_DrawBalloon_DefaultColors(t *testing.T) {
	// Test balloon with nil colors (use defaults)
	opts := Options{
		Size:            32,
		BaseSize:        32,
		Multiple:        1,
		Style:           "basic",
		Theme:           "white",
		BackgroundColor: &color.RGBA{R: 0, G: 0, B: 0, A: 255},
		Format:          "png",
		Output:          io.Discard,
		Text:            "Test",
		Balloon:         true,
	}
	p := NewPortrait(opts)
	err := p.Encode()
	assert.NoError(t, err, "Portrait.Encode() with balloon (nil colors) should not fail")
}

func TestPortrait_DrawBalloon_WithTextColor(t *testing.T) {
	// Test balloon with explicit text color
	opts := Options{
		Size:            64,
		BaseSize:        32,
		Multiple:        2,
		Style:           "basic",
		Theme:           "white",
		BackgroundColor: &color.RGBA{R: 26, G: 26, B: 26, A: 255},
		Format:          "png",
		Output:          io.Discard,
		Text:            "Cock-a-doodle-doo!",
		TextColor:       &color.RGBA{R: 255, G: 0, B: 0, A: 255},
		TextFontSize:    14,
		Balloon:         true,
	}
	p := NewPortrait(opts)
	err := p.Encode()
	assert.NoError(t, err, "Portrait.Encode() with balloon and explicit text color should not fail")
}

func TestPortrait_Balloon_NotUsedWithoutText(t *testing.T) {
	// Balloon flag without text should not affect output (no balloon rendered)
	opts := Options{
		Size:            32,
		BaseSize:        32,
		Multiple:        1,
		Style:           "basic",
		Theme:           "white",
		BackgroundColor: &color.RGBA{R: 0, G: 0, B: 0, A: 255},
		Format:          "png",
		Output:          io.Discard,
		Balloon:         true,
		// Text is empty
	}
	p := NewPortrait(opts)
	err := p.Encode()
	assert.NoError(t, err, "Portrait.Encode() with balloon but no text should not fail")
}
