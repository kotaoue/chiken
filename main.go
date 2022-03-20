package main

import (
	"errors"
	"flag"
	"fmt"
	"image/color"
	"os"
	"strings"

	"github.com/kotaoue/chiken/pkg/cutil"
	"github.com/kotaoue/chiken/pkg/portrait"
)

var (
	theme      = flag.String("t", portrait.WhiteTheme, "theme color of rooster")
	style      = flag.String("s", portrait.BasicStyle, "style of rooster")
	multiple   = flag.Int("m", 1, "value to be multiplied by 32")
	format     = flag.String("f", "png", "format of output image")
	delay      = flag.Int("d", 0, "delay time for gif")
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
	return output()
}

func output() error {
	if err := checkFormat(*format); err != nil {
		return err
	}

	c, err := cutil.HexToColor(*background)
	if err != nil {
		return err
	}

	if err := encode(c); err != nil {
		return err
	}

	printReference()

	return nil
}

func encode(c *color.RGBA) error {
	p, err := portrait.NewPortrait(
		portrait.Options{
			Size:            size,
			BaseSize:        baseSize,
			Multiple:        *multiple,
			Style:           *style,
			Theme:           *theme,
			BackgroundColor: c,
			Format:          *format,
			Delay:           *delay,
			FileName:        fileName(),
		},
	)

	if err != nil {
		return err
	}

	return p.Encode()
}

func printReference() {
	alt := fileName()
	alt = strings.TrimPrefix(alt, "img/")
	alt = strings.TrimSuffix(alt, fmt.Sprintf(".%s", *format))

	fmt.Printf(
		"|%s|%s|%s|%d*%d|%s|![%s](%s)|\n",
		strings.Join(os.Args[1:], " "),
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

	if *style != portrait.BasicStyle {
		name = fmt.Sprintf("%s_%s", name, *style)
	}
	if *multiple > 1 {
		name = fmt.Sprintf("%s_%d", name, *multiple)
	}
	if *background != "transparent" {
		name = fmt.Sprintf("%s_%s", name, strings.ReplaceAll(*background, "#", ""))
	}
	if *delay > 0 {
		name = fmt.Sprintf("%s_delay%d", name, *delay)
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
