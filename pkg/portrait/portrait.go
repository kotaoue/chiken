package portrait

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

type Portrait struct {
	opt Options
}

type Options struct {
	Size            int
	BaseSize        int
	Multiple        int
	Style           string
	Theme           string
	BackgroundColor *color.RGBA
	FileName        string
}

func NewPortrait(o Options) *Portrait {
	return &Portrait{opt: o}
}

func (p *Portrait) Draw() error {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{p.opt.Size, p.opt.Size}})

	p.drawBackground(img)

	if err := p.drawSubject(img); err != nil {
		return err
	}

	f, err := os.Create(p.opt.FileName)
	if err != nil {
		return err
	}

	return png.Encode(f, img)
}

func (p *Portrait) drawBackground(img *image.RGBA) {
	for x := 0; x < p.opt.Size; x++ {
		for y := 0; y < p.opt.Size; y++ {
			img.Set(x, y, p.opt.BackgroundColor)
		}
	}
}

func (p *Portrait) drawSubject(img *image.RGBA) error {
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
