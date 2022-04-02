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
	opt Options
}

type Options struct {
	Size            int
	BaseSize        int
	Multiple        int
	Style           string
	Theme           string
	Effect          string
	BackgroundColor *color.RGBA
	Format          string
	Delay           int
	FileName        string
}

func NewPortrait(o Options) *Portrait {
	return &Portrait{opt: o}
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
	img, err := p.draw(p.opt.Style, p.opt.Theme, p.opt.Effect)
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
	thms, err := Theme{}.Adjust(p.opt.Theme)
	if err != nil {
		return err
	}
	fmt.Println(thms)

	var images []*image.Paletted
	var delays []int
	var disposals []byte

	for _, thm := range strings.Split(thms, "-") {
		for _, stl := range strings.Split(p.opt.Style, "-") {
			fmt.Printf("style:%s theme:%s\n", stl, thm)
			img, err := p.draw(stl, thm, p.opt.Effect)
			if err != nil {
				return err
			}

			images = append(images, img)
			delays = append(delays, p.opt.Delay)
			disposals = append(disposals, gif.DisposalPrevious)
		}
	}

	fp, err := os.OpenFile(p.opt.FileName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer fp.Close()

	return gif.EncodeAll(fp, &gif.GIF{
		Image:    images,
		Delay:    delays,
		Disposal: disposals,
	})
}

func (p *Portrait) draw(stl string, thm string, eff string) (*image.Paletted, error) {
	style, err := p.fetchStyle(stl)
	if err != nil {
		return nil, err
	}

	theme, err := p.fetchTheme(thm)
	if err != nil {
		return nil, err
	}

	effect := NewEffect(theme)
	theme, err = effect.Apply(eff)
	if err != nil {
		return nil, err
	}

	img := image.NewPaletted(image.Rectangle{image.Point{0, 0}, image.Point{p.opt.Size, p.opt.Size}}, theme)

	p.drawBackground(img)

	if err := p.drawSubject(img, style, theme); err != nil {
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

func (p *Portrait) drawSubject(img *image.Paletted, subject [][]int, theme []color.Color) error {
	for y := 0; y < p.opt.BaseSize; y++ {
		fmt.Print("{")
		for x := 0; x < p.opt.BaseSize; x++ {
			i := subject[y][x]
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
