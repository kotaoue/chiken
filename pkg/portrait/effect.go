package portrait

import (
	"errors"
	"fmt"
	"image/color"
	"strconv"
	"strings"
)

const (
	Mirror    = "mirror"
	Negative  = "negative"
	Grayscale = "grayscale"
	RightLoop = "rightLoop"
	LeftLoop  = "leftLoop"
)

type Effect struct {
	style [][]int
	theme []color.Color
}

func NewEffect(stl [][]int, thm []color.Color, size int) *Effect {
	return &Effect{
		style: stl,
		theme: thm,
	}
}

func (e Effect) Adjust(effects string, size int) (string, error) {
	switch {
	case strings.HasPrefix(effects, RightLoop):
		return e.splitLoop(RightLoop, effects, size)
	case strings.HasPrefix(effects, LeftLoop):
		return e.splitLoop(LeftLoop, effects, size)
	}

	return effects, nil
}

func (e Effect) splitLoop(prefix, effects string, size int) (string, error) {
	step, err := strconv.Atoi(strings.TrimPrefix(effects, prefix))
	if err != nil || step <= 0 {
		return "", fmt.Errorf("effects expects a format of %s{int}. but '%s' was specified", effects, prefix)
	}

	var ss []string
	for i := 0; i < size; i += step {
		ss = append(ss, fmt.Sprintf("%s%d", prefix, i))
	}
	return strings.Join(ss, "-"), nil
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
	switch {
	case strings.HasPrefix(effect, RightLoop):
		return e.rightLoop(effect)
	case strings.HasPrefix(effect, LeftLoop):
		return e.leftLoop(effect)
	}

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

func (e Effect) rightLoop(s string) error {
	step, err := strconv.Atoi(strings.TrimPrefix(s, RightLoop))
	if err != nil || step < 0 {
		return fmt.Errorf("effect expects a format of rightLoop{int}. but '%s' was specified", s)
	}

	for y, dots := range e.style {
		row := make([]int, len(dots))

		for i := 0; i < len(dots); i++ {
			row[(i+step)%(len(dots)-1)] = dots[i]
		}
		e.style[y] = row
	}
	return nil
}

func (e Effect) leftLoop(s string) error {
	step, err := strconv.Atoi(strings.TrimPrefix(s, LeftLoop))
	if err != nil || step < 0 {
		return fmt.Errorf("effect expects a format of leftLoop{int}. but '%s' was specified", s)
	}

	for y, dots := range e.style {
		row := make([]int, len(dots))

		for i := 0; i < len(dots); i++ {
			row[i] = dots[(i+step)%(len(dots)-1)]
		}
		e.style[y] = row
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
