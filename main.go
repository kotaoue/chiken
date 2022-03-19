package main

import (
	"errors"
	"flag"
	"fmt"
	"image/color"
	"os"
	"strings"

	"github.com/kotaoue/chiken/pkg/portrait"
	ptr "github.com/kotaoue/chiken/pkg/portrait"
)

var (
	theme      = flag.String("t", ptr.WhiteTheme, "theme color of rooster")
	style      = flag.String("s", ptr.BasicStyle, "style of rooster")
	multiple   = flag.Int("m", 1, "value to be multiplied by 32")
	format     = flag.String("f", "png", "format of output image")
	background = flag.String("b", "transparent", "background color. set with hex. example #ffffff. empty is transparent")
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
	printReference()

	return output(c)
}

func output(c *color.RGBA) error {
	p := portrait.NewPortrait(
		portrait.Options{
			Size:            size,
			BaseSize:        baseSize,
			Multiple:        *multiple,
			Style:           *style,
			Theme:           *theme,
			BackgroundColor: c,
			FileName:        fileName(),
		},
	)

	return p.Draw()
}

func printReference() {
	args := strings.Join(os.Args[1:], " ")
	if args != "" {
		args = " " + args
	}

	alt := fileName()
	alt = strings.TrimPrefix(alt, "img/")
	alt = strings.TrimSuffix(alt, fmt.Sprintf(".%s", *format))

	fmt.Printf(
		"|go run main.go%s|%s|%s|%d*%d|%s|![%s](%s)|\n",
		args,
		*theme,
		*style,
		size,
		size,
		*background,
		alt,
		fileName(),
	)
}

func fileName() string {
	dir := "img"
	name := *theme

	if *style != ptr.BasicStyle {
		name = fmt.Sprintf("%s_%s", name, *style)
	}
	if *multiple > 1 {
		name = fmt.Sprintf("%s_%d", name, *multiple)
	}
	if *background != "transparent" {
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
