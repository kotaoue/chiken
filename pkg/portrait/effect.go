package portrait

import (
	"errors"
	"image/color"
	"strings"
)

const (
	Negative  = "negative"
	Grayscale = "grayscale"
)

type Effect struct {
	theme []color.Color
}

func NewEffect(thm []color.Color) *Effect {
	return &Effect{theme: thm}
}

func (e Effect) Apply(effects string) ([]color.Color, error) {
	for _, effect := range strings.Split(effects, "-") {
		theme, err := e.apply(effect)
		if err != nil {
			return nil, err
		}

		e.theme = theme
	}

	return e.theme, nil
}

func (e Effect) apply(effect string) ([]color.Color, error) {
	if effect != "" {
		switch effect {
		case Negative:
			e.negative()
		case Grayscale:
			e.grayscale()
		default:
			return nil, errors.New("not exist effect")
		}
	}
	return e.theme, nil
}

func (e Effect) negative() []color.Color {
	for k, v := range e.theme {
		r, g, b, a := v.RGBA()
		e.theme[k] = color.RGBA{
			R: uint8(^r),
			G: uint8(^g),
			B: uint8(^b),
			A: uint8(a),
		}
		vPrintf("%3v -> %3v\n", v, e.theme[k])
	}
	return e.theme
}

func (e Effect) grayscale() []color.Color {
	for k, v := range e.theme {
		r, g, b, a := v.RGBA()
		gray := float64(r)*0.3 + float64(g)*0.59 + float64(b)*0.11
		e.theme[k] = color.RGBA{
			R: uint8(gray),
			G: uint8(gray),
			B: uint8(gray),
			A: uint8(a),
		}
		vPrintf("%3v -> %3v\n", v, e.theme[k])
	}
	return e.theme
}
