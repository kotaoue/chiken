package portrait

import (
	"fmt"
	"image"
	"image/color"
)

type Portrait struct {
	opt Options
}

type Options struct {
	Size     int
	BaseSize int
	Multiple int
	Style    string
	Theme    string
}

func NewPortrait(o Options) *Portrait {
	return &Portrait{opt: o}
}

func (p *Portrait) BG(img *image.RGBA, col *color.RGBA) {
	for x := 0; x < p.opt.Size; x++ {
		for y := 0; y < p.opt.Size; y++ {
			img.Set(x, y, col)
		}
	}
}

func (p *Portrait) Draw(img *image.RGBA) error {
	style, theme, err := p.fetchBlueprint()
	if err != nil {
		return err
	}

	for y := 0; y < p.opt.BaseSize; y++ {
		fmt.Print("{")
		for x := 0; x < p.opt.BaseSize; x++ {
			i := style[y][x]
			fmt.Print(i)

			if i != 0 {
				for my := 0; my < p.opt.Multiple; my++ {
					for mx := 0; mx < p.opt.Multiple; mx++ {
						img.Set(x*p.opt.Multiple+mx, y*p.opt.Multiple+my, theme[i])
					}
				}
			}
		}
		fmt.Println("}")
	}

	return nil
}

func (p *Portrait) fetchBlueprint() ([][]int, []color.RGBA, error) {
	theme, err := Theme{}.Get(p.opt.Theme)
	if err != nil {
		return nil, nil, err
	}

	return Style{}.Get(p.opt.Style), theme, nil
}
