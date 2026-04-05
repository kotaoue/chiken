package portrait

import (
	"image/color"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateBalloonGrid(t *testing.T) {
	// Test with border=1 (minimum)
	border := 1
	innerW := 20
	innerH := 10
	grid := generateBalloonGrid(innerW, innerH, border)

	bodyW := innerW + border*2
	bodyH := innerH + border*2
	tailH := border * 4
	totalH := bodyH + tailH

	assert.Equal(t, totalH, len(grid), "grid height should match bodyH + tailH")
	for _, row := range grid {
		assert.Equal(t, bodyW, len(row), "each grid row width should match bodyW")
	}

	// Corners should be transparent (0)
	cornerCut := border * 2
	assert.Equal(t, 0, grid[0][0], "top-left corner should be transparent")
	assert.Equal(t, 0, grid[0][bodyW-1], "top-right corner should be transparent")
	assert.Equal(t, 0, grid[bodyH-1][0], "bottom-left corner should be transparent")
	assert.Equal(t, 0, grid[bodyH-1][bodyW-1], "bottom-right corner should be transparent")

	// Inside corner cut area should be transparent
	if cornerCut > 1 {
		assert.Equal(t, 0, grid[0][1], "top-left inner corner should be transparent")
	}

	// Body border cells (non-corner) should be 1
	assert.Equal(t, 1, grid[0][cornerCut], "top border should be 1")
	assert.Equal(t, 1, grid[cornerCut][0], "left border should be 1")

	// Inner body cells should be fill (2)
	assert.Equal(t, 2, grid[cornerCut][cornerCut], "inner body cell should be fill (2)")

	// Tail cells should be present (1 or 2) in tail area
	tailX := cornerCut
	tailFound := false
	for y := bodyH; y < totalH; y++ {
		for x := tailX; x < tailX+border*3 && x < bodyW; x++ {
			if grid[y][x] != 0 {
				tailFound = true
			}
		}
	}
	assert.True(t, tailFound, "tail cells should be present")
}

func TestGenerateBalloonGrid_Border2(t *testing.T) {
	// Test with border=2 (2x multiple)
	border := 2
	innerW := 40
	innerH := 20
	grid := generateBalloonGrid(innerW, innerH, border)

	bodyW := innerW + border*2
	bodyH := innerH + border*2
	tailH := border * 4
	totalH := bodyH + tailH

	assert.Equal(t, totalH, len(grid), "grid height should match bodyH + tailH")
	for _, row := range grid {
		assert.Equal(t, bodyW, len(row), "each grid row width should match bodyW")
	}

	// Corners should be transparent
	cornerCut := border * 2
	assert.Equal(t, 0, grid[0][0], "top-left corner should be transparent")
	assert.Equal(t, 0, grid[0][cornerCut-1], "top-left inner corner should be transparent")

	// Non-corner top border should be 1
	assert.Equal(t, 1, grid[0][cornerCut], "top border should be 1")

	// Inner body cells should be fill (2) - use cornerCut offset to avoid corner-cut area
	cornerCut2 := border * 2
	assert.Equal(t, 2, grid[cornerCut2][cornerCut2], "inner body cell should be fill (2)")
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
