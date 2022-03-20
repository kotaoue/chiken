package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"image/color"
	"os"
	"strconv"
	"strings"

	"github.com/kotaoue/chiken/pkg/cutil"
	"github.com/kotaoue/chiken/pkg/portrait"
)

var (
	theme      = flag.String("t", portrait.WhiteTheme, "theme color of rooster")
	style      = flag.String("s", portrait.BasicStyle, "style of rooster")
	format     = flag.String("f", "png", "format of output image")
	background = flag.String("b", "transparent", "background color. set with hex. example #ffffff. empty is transparent")
	multiple   = flag.Int("m", 1, "value to be multiplied by 32")
	delay      = flag.Int("d", 0, "delay time for gif")
	dump       = flag.Bool("dump", false, "re encode from Args Example on README")
	size       int
	baseSize   int
)

func init() {
	flag.Parse()

	baseSize = 32
}

func main() {
	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Main() error {
	if *dump {
		return reOutputs()
	}
	return output()
}

func output() error {
	size = baseSize * *multiple

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

	args := strings.Join(os.Args[1:], " ")

	fmt.Printf(
		"|%s|%s|%s|%d*%d|%s|![%s](%s)|\n",
		strings.ReplaceAll(args, "-dump", ""),
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

func reOutputs() error {
	file, err := os.Open("README.md")
	if err != nil {
		return err
	}

	fs := bufio.NewScanner(file)
	afterArgsLine := false
	afterHyphenLine := false
	for fs.Scan() {
		if afterArgsLine && afterHyphenLine {
			ss := strings.Split(fs.Text(), "|")

			*theme = portrait.WhiteTheme
			*style = portrait.BasicStyle
			*format = "png"
			*background = "transparent"
			*multiple = 1
			*delay = 0

			for _, v := range strings.Split(ss[1], " ") {
				switch {
				case strings.HasPrefix(v, "-s="):
					*style = strings.TrimPrefix(v, "-s=")
				case strings.HasPrefix(v, "-t="):
					*theme = strings.TrimPrefix(v, "-t=")
				case strings.HasPrefix(v, "-f="):
					*format = strings.TrimPrefix(v, "-f=")
				case strings.HasPrefix(v, "-b="):
					*background = strings.TrimPrefix(v, "-b=")
				case strings.HasPrefix(v, "-m="):
					i, err := strconv.Atoi(strings.TrimPrefix(v, "-m="))
					if err != nil {
						return err
					}
					*multiple = i
				case strings.HasPrefix(v, "-d="):
					i, err := strconv.Atoi(strings.TrimPrefix(v, "-d="))
					if err != nil {
						return err
					}
					*delay = i
				}
			}

			if err := output(); err != nil {
				return err
			}
		}

		switch fs.Text() {
		case "## Args Example":
			afterArgsLine = true
		case "|---|---|---|---|---|---|":
			afterHyphenLine = true
		}
	}

	return nil
}
