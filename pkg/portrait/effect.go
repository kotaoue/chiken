package portrait

import (
	"errors"
	"fmt"
	"image/color"
)

const (
	Negative = "negative"
)

type Effect struct {
	theme []color.Color
}

func NewEffect(thm []color.Color) *Effect {
	return &Effect{theme: thm}
}

func (e Effect) Apply(effect string) ([]color.Color, error) {
	if effect != "" {
		switch effect {
		case Negative:
			e.negative()
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
		fmt.Printf("%3v -> %3v\n", v, e.theme[k])
	}
	return e.theme
}
