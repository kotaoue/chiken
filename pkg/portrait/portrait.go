package portrait

import (
	"image"
	"image/color"
)

type Portrait struct {
	size int
}

func NewPortrait(size int) *Portrait {
	return &Portrait{
		size: size,
	}
}

func (d *Portrait) BG(img *image.RGBA, col *color.RGBA) {
	for x := 0; x < d.size; x++ {
		for y := 0; y < d.size; y++ {
			img.Set(x, y, col)
		}
	}
}
