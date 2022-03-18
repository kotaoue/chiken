package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"

	"github.com/kotaoue/chiken/pkg/blueprint"
	"github.com/kotaoue/chiken/pkg/palette"
	"github.com/kotaoue/chiken/pkg/portrait"
)

var (
	theme      = flag.Int("t", 0, "theme color of rooster")
	style      = flag.Int("s", 0, "style of rooster")
	multiple   = flag.Int("m", 1, "value to be multiplied by 32")
	format     = flag.String("f", "png", "format of output image")
	background = flag.String("b", "", "background color. set with hex. example #ffffff. empty is transparent")
	size       int
	baseSize   int
)

func init() {
	flag.Parse()

	baseSize = 32
	size = baseSize * *multiple
}

func main() {
	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Main() error {
	if err := checkFormat(*format); err != nil {
		return err
	}

	c, err := hexToColor(*background)
	if err != nil {
		return err
	}
	fmt.Printf("size:%d multiple:%d background:%v\n", size, *multiple, c)

	return output(c)
}

func output(c *color.RGBA) error {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{size, size}})

	p := portrait.NewPortrait(
		portrait.Options{
			Size:     size,
			BaseSize: baseSize,
			Multiple: *multiple,
			Style:    *style,
			Theme:    *theme,
		},
	)
	p.BG(img, c)

	if err := p.Draw(img); err != nil {
		return err
	}

	f, err := os.Create(fileName())
	if err != nil {
		return err
	}

	return png.Encode(f, img)
}

func fileName() string {
	dir := "img"
	name := palette.Name(*theme)

	if *style != blueprint.BasicStyle {
		name = fmt.Sprintf("%s_%s", name, blueprint.Name(*style))
	}
	if *multiple > 1 {
		name = fmt.Sprintf("%s_%d", name, *multiple)
	}
	if *background != "" {
		name = fmt.Sprintf("%s_%s", name, strings.ReplaceAll(*background, "#", ""))
	}
	return fmt.Sprintf("%s/%s.%s", dir, name, *format)
}

func checkFormat(s string) error {
	switch s {
	case "gif", "png":
		return nil
	}

	return errors.New("Unsupported formats")
}

func hexToColor(s string) (*color.RGBA, error) {
	c := &color.RGBA{}

	if len(s) == 7 {
		if _, err := fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B); err != nil {
			return nil, err
		}
		c.A = 255
	}

	return c, nil
}
