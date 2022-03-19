package portrait

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"os"
	"strings"
)

type Portrait struct {
	opt   Options
	theme []color.Color
}

type Options struct {
	Size            int
	BaseSize        int
	Multiple        int
	Style           string
	Theme           string
	BackgroundColor *color.RGBA
	Format          string
	Delay           int
	FileName        string
}

func NewPortrait(o Options) (*Portrait, error) {
	p := &Portrait{opt: o}

	theme, err := p.fetchTheme(o.Theme)
	if err != nil {
		return nil, err
	}

	p.theme = theme
	return p, err
}

func (p *Portrait) Encode() error {
	switch p.opt.Format {
	case "png":
		return p.encodePng()
	case "gif":
		return p.encodeGif()
	}
	return nil
}

func (p *Portrait) encodePng() error {
	img, err := p.draw(p.opt.Style)
	if err != nil {
		return err
	}

	f, err := os.Create(p.opt.FileName)
	if err != nil {
		return err
	}

	return png.Encode(f, img)
}

func (p *Portrait) encodeGif() error {
	var images []*image.Paletted
	var delays []int

	for _, v := range strings.Split(p.opt.Style, "-") {
		fmt.Println(v)
		img, err := p.draw(v)
		if err != nil {
			return err
		}

		images = append(images, img)
		delays = append(delays, p.opt.Delay)
	}

	fp, err := os.OpenFile(p.opt.FileName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer fp.Close()

	return gif.EncodeAll(fp, &gif.GIF{
		Image: images,
		Delay: delays,
	})
}

func (p *Portrait) draw(s string) (*image.Paletted, error) {
	style, err := p.fetchStyle(s)
	if err != nil {
		return nil, err
	}
	img := image.NewPaletted(image.Rectangle{image.Point{0, 0}, image.Point{p.opt.Size, p.opt.Size}}, p.theme)

	p.drawBackground(img)

	if err := p.drawSubject(img, style); err != nil {
		return nil, err
	}
	return img, nil
}

func (p *Portrait) drawBackground(img *image.Paletted) {
	for x := 0; x < p.opt.Size; x++ {
		for y := 0; y < p.opt.Size; y++ {
			img.Set(x, y, p.opt.BackgroundColor)
		}
	}
}

func (p *Portrait) drawSubject(img *image.Paletted, subject [][]int) error {
	for y := 0; y < p.opt.BaseSize; y++ {
		fmt.Print("{")
		for x := 0; x < p.opt.BaseSize; x++ {
			i := subject[y][x]
			fmt.Print(i)

			if i != 0 {
				for my := 0; my < p.opt.Multiple; my++ {
					for mx := 0; mx < p.opt.Multiple; mx++ {
						img.Set(x*p.opt.Multiple+mx, y*p.opt.Multiple+my, p.theme[i])
					}
				}
			}
		}
		fmt.Println("}")
	}

	return nil
}

func (p *Portrait) fetchStyle(s string) ([][]int, error) {
	style, err := Style{}.Get(s)
	if err != nil {
		return nil, err
	}

	return style, nil
}

func (*Portrait) fetchTheme(s string) ([]color.Color, error) {
	theme, err := Theme{}.Get(s)
	if err != nil {
		return nil, err
	}

	return theme, nil
}
