package portrait

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/png"
	"io"
	"os"
	"strings"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
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
	Verbose         bool
	Output          io.Writer
	Text            string
	TextColor       *color.RGBA
	TextFontSize    int
}

func NewPortrait(o Options) *Portrait {
	verbose = o.Verbose
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

	var finalImg image.Image = img
	if p.opt.Text != "" {
		finalImg = p.drawText(img)
	}

	w := p.opt.Output
	if w == nil {
		f, err := os.Create(p.opt.FileName)
		if err != nil {
			return err
		}
		defer f.Close()
		w = f
	}

	return png.Encode(w, finalImg)
}

func (p *Portrait) encodeGif() error {
	thms, err := Theme{}.Adjust(p.opt.Theme)
	if err != nil {
		return err
	}

	eff := NewEffect([][]int{}, []color.Color{}, p.opt.BaseSize)
	effs, err := eff.Adjust(p.opt.Effect, p.opt.BaseSize)
	if err != nil {
		return err
	}

	var images []*image.Paletted
	var delays []int
	var disposals []byte

	for _, eff := range strings.Split(effs, "-") {
		for _, thm := range strings.Split(thms, "-") {
			for _, stl := range strings.Split(p.opt.Style, "-") {
				vPrintf("style:%s theme:%s\n", stl, thm)
				img, err := p.draw(stl, thm, eff)
				if err != nil {
					return err
				}

				images = append(images, img)
				delays = append(delays, p.opt.Delay)
				disposals = append(disposals, gif.DisposalPrevious)
			}
		}
	}

	w := p.opt.Output
	if w == nil {
		fp, err := os.OpenFile(p.opt.FileName, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		defer fp.Close()
		w = fp
	}

	return gif.EncodeAll(w, &gif.GIF{
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
	theme = append(theme, *p.opt.BackgroundColor)

	effect := NewEffect(style, theme, p.opt.BaseSize)
	style, theme, err = effect.Apply(eff)
	if err != nil {
		return nil, err
	}

	img := image.NewPaletted(image.Rectangle{image.Point{0, 0}, image.Point{p.opt.Size, p.opt.Size}}, theme)

	p.drawBackground(img, theme)

	if err := p.drawSubject(img, style, theme); err != nil {
		return nil, err
	}
	return img, nil
}

func (p *Portrait) drawBackground(img *image.Paletted, theme []color.Color) {
	for x := 0; x < p.opt.Size; x++ {
		for y := 0; y < p.opt.Size; y++ {
			img.Set(
				x,
				y,
				theme[len(theme)-1],
			)
		}
	}
}

func (p *Portrait) drawSubject(img *image.Paletted, subject [][]int, theme []color.Color) error {
	for y := 0; y < p.opt.BaseSize; y++ {
		vPrint("{")
		for x := 0; x < p.opt.BaseSize; x++ {
			i := subject[y][x]
			vPrint(i)

			if i != 0 {
				for my := 0; my < p.opt.Multiple; my++ {
					for mx := 0; mx < p.opt.Multiple; mx++ {
						img.Set(x*p.opt.Multiple+mx, y*p.opt.Multiple+my, theme[i])
					}
				}
			}
		}
		vPrintln("}")
	}

	return nil
}

func (p *Portrait) newFontFace() font.Face {
	if p.opt.TextFontSize > 0 {
		tt, err := opentype.Parse(goregular.TTF)
		if err != nil {
			vPrintf("failed to parse font: %v\n", err)
			return basicfont.Face7x13
		}
		f, err := opentype.NewFace(tt, &opentype.FaceOptions{
			Size: float64(p.opt.TextFontSize),
			DPI:  72,
		})
		if err != nil {
			vPrintf("failed to create font face: %v\n", err)
			return basicfont.Face7x13
		}
		return f
	}
	return basicfont.Face7x13
}

func (p *Portrait) drawText(portrait *image.Paletted) image.Image {
	face := p.newFontFace()
	if closer, ok := face.(io.Closer); ok {
		defer closer.Close()
	}

	textColor := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	if p.opt.TextColor != nil && p.opt.TextColor.A != 0 {
		textColor = *p.opt.TextColor
	}

	padding := p.opt.Size / 4
	textWidth := font.MeasureString(face, p.opt.Text).Ceil()
	canvasWidth := p.opt.Size + padding + textWidth + padding

	canvas := image.NewNRGBA(image.Rect(0, 0, canvasWidth, p.opt.Size))

	draw.Draw(canvas, canvas.Bounds(), image.NewUniform(p.opt.BackgroundColor), image.Point{}, draw.Src)
	draw.Draw(canvas, portrait.Bounds(), portrait, image.Point{}, draw.Over)

	metrics := face.Metrics()
	textY := (p.opt.Size + metrics.Ascent.Ceil() - metrics.Descent.Ceil()) / 2

	d := &font.Drawer{
		Dst:  canvas,
		Src:  image.NewUniform(textColor),
		Face: face,
		Dot:  fixed.P(p.opt.Size+padding, textY),
	}
	d.DrawString(p.opt.Text)

	return canvas
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
