package portrait

import (
	"errors"
	"image/color"
	"strings"
)

const (
	Mirror    = "mirror"
	Negative  = "negative"
	Grayscale = "grayscale"
)

type Effect struct {
	style [][]int
	theme []color.Color
}

func NewEffect(stl [][]int, thm []color.Color) *Effect {
	return &Effect{
		style: stl,
		theme: thm,
	}
}

func (e Effect) Apply(effects string) ([][]int, []color.Color, error) {
	for _, effect := range strings.Split(effects, "-") {
		if err := e.apply(effect); err != nil {
			return nil, nil, err
		}
	}

	return e.style, e.theme, nil
}

func (e Effect) apply(effect string) error {
	if effect != "" {
		switch effect {
		case Mirror:
			e.mirror()
		case Negative:
			e.negative()
		case Grayscale:
			e.grayscale()
		default:
			return errors.New("not exist effect")
		}
	}
	return nil
}

func (e Effect) mirror() {
	for y, dots := range e.style {
		row := make([]int, len(dots))

		for i := 0; i < len(dots); i++ {
			row[i] = dots[len(dots)-i-1]
		}
		e.style[y] = row
	}
}

func (e Effect) negative() {
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
}

func (e Effect) grayscale() {
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
}
