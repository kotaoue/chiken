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
	UpLoop    = "upLoop"
	DownLoop  = "downLoop"
	RotateClockwise = "rotateClockwise"
	RotateCounterClockwise = "rotateCounterClockwise"
	
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

func (e *Effect) Adjust(effects string, size int) (string, error) {
	switch {
	case strings.HasPrefix(effects, RightLoop):
		return e.splitLoop(RightLoop, effects, size)
	case strings.HasPrefix(effects, LeftLoop):
		return e.splitLoop(LeftLoop, effects, size)
	case strings.HasPrefix(effects, UpLoop):
		return e.splitLoop(UpLoop, effects, size)
	case strings.HasPrefix(effects, DownLoop):
		return e.splitLoop(DownLoop, effects, size)
	}

	return effects, nil
}

func (e *Effect) splitLoop(prefix, effects string, size int) (string, error) {
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

func (e *Effect) Apply(effects string) ([][]int, []color.Color, error) {
	for _, effect := range strings.Split(effects, "-") {
		if err := e.apply(effect); err != nil {
			return nil, nil, err
		}
	}

	return e.style, e.theme, nil
}

func (e *Effect) apply(effect string) error {
	switch {
	case strings.HasPrefix(effect, RightLoop):
		return e.rightLoop(effect)
	case strings.HasPrefix(effect, LeftLoop):
		return e.leftLoop(effect)
	case strings.HasPrefix(effect, UpLoop):
		return e.upLoop(effect)
	case strings.HasPrefix(effect, DownLoop):
		return e.downLoop(effect)
	}

	if effect != "" {
		switch effect {
		case Mirror:
			e.mirror()
		case Negative:
			e.negative()
		case Grayscale:
			e.grayscale()
		case RotateClockwise:
			e.rotateClockwise()
		case RotateCounterClockwise:
			e.rotateCounterClockwise()
		default:
			return errors.New("not exist effect")
		}
	}
	return nil
}

func (e *Effect) rightLoop(s string) error {
	step, err := strconv.Atoi(strings.TrimPrefix(s, RightLoop))
	if err != nil || step < 0 {
		return fmt.Errorf("effect expects a format of rightLoop{int}. but '%s' was specified", s)
	}

	for y, dots := range e.style {
		row := make([]int, len(dots))

		for i := 0; i < len(dots); i++ {
			row[(i+step)%len(dots)] = dots[i]
		}
		e.style[y] = row
	}
	return nil
}

func (e *Effect) leftLoop(s string) error {
	step, err := strconv.Atoi(strings.TrimPrefix(s, LeftLoop))
	if err != nil || step < 0 {
		return fmt.Errorf("effect expects a format of leftLoop{int}. but '%s' was specified", s)
	}

	for y, dots := range e.style {
		row := make([]int, len(dots))

		for i := 0; i < len(dots); i++ {
			row[i] = dots[(i+step)%len(dots)]
		}
		e.style[y] = row
	}
	return nil
}

func (e *Effect) upLoop(s string) error {
	step, err := strconv.Atoi(strings.TrimPrefix(s, UpLoop))
	if err != nil || step < 0 {
		return fmt.Errorf("effect expects a format of upLoop{int}. but '%s' was specified", s)
	}

	style := make([][]int, len(e.style))
	for y := 0; y < len(style); y++ {
		style[y] = e.style[(y+step)%(len(e.style))]
	}

	copy(e.style, style)
	return nil
}

func (e *Effect) downLoop(s string) error {
	step, err := strconv.Atoi(strings.TrimPrefix(s, DownLoop))
	if err != nil || step < 0 {
		return fmt.Errorf("effect expects a format of upLoop{int}. but '%s' was specified", s)
	}

	style := make([][]int, len(e.style))
	for y := 0; y < len(style); y++ {
		style[(y+step)%(len(e.style))] = e.style[y]
	}

	copy(e.style, style)
	return nil
}

func (e *Effect) mirror() {
	for y, dots := range e.style {
		row := make([]int, len(dots))

		for i := 0; i < len(dots); i++ {
			row[i] = dots[len(dots)-i-1]
		}
		e.style[y] = row
	}
}

func (e *Effect) negative() {
	for k, v := range e.theme {
		if c, ok := v.(color.RGBA); ok {
			e.theme[k] = color.RGBA{
				R: ^c.R,
				G: ^c.G,
				B: ^c.B,
				A: c.A,
			}
			vPrintf("%3v -> %3v\n", v, e.theme[k])
		}
	}
}

func (e *Effect) grayscale() {
	for k, v := range e.theme {
		if c, ok := v.(color.RGBA); ok {
			gray := float64(c.R)*0.3 + float64(c.G)*0.59 + float64(c.B)*0.11
			e.theme[k] = color.RGBA{
				R: uint8(gray),
				G: uint8(gray),
				B: uint8(gray),
				A: c.A,
			}
			vPrintf("%3v -> %3v\n", v, e.theme[k])
		}
	}
}

func (e *Effect) rotateClockwise() {
    style := make([][]int, len(e.style[0]))
    for i := range style {
        style[i] = make([]int, len(e.style))
    }

    for y := 0; y < len(e.style); y++ {
        for x := 0; x < len(e.style[0]); x++ {
            style[x][len(e.style)-1-y] = e.style[y][x]
        }
    }

    e.style = style
}

func (e *Effect) rotateCounterClockwise() {
	style := make([][]int, len(e.style[0]))
	for i := range style {
		style[i] = make([]int, len(e.style))
	}

	for y := 0; y < len(e.style); y++ {
		for x := 0; x < len(e.style[0]); x++ {
			style[len(e.style[0])-1-x][y] = e.style[y][x]
		}
	}

	e.style = style
}
