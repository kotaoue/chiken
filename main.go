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

const (
	defaultTheme      = portrait.WhiteTheme
	defaultStyle      = portrait.BasicStyle
	defaultFormat     = "png"
	defaultEffect     = ""
	defaultBackground = "transparent"
	defaultMultiple   = 1
	defaultDelay      = 0
)

var (
	theme      = flag.String("t", defaultTheme, "theme color of rooster")
	style      = flag.String("s", defaultStyle, "style of rooster")
	format     = flag.String("f", defaultFormat, "format of output image")
	effect     = flag.String("e", defaultEffect, "set visual effects")
	background = flag.String("b", defaultBackground, "background color. set with hex. example #ffffff. empty is transparent")
	multiple   = flag.Int("m", defaultMultiple, "value to be multiplied by 32")
	delay      = flag.Int("d", defaultDelay, "delay time for gif")
	verbose    = flag.Bool("v", false, "printing verbose output")
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
	p := portrait.NewPortrait(
		portrait.Options{
			Size:            size,
			BaseSize:        baseSize,
			Multiple:        *multiple,
			Style:           *style,
			Theme:           *theme,
			BackgroundColor: c,
			Format:          *format,
			Effect:          *effect,
			Delay:           *delay,
			FileName:        fileName(),
			Verbose:         *verbose,
		},
	)

	return p.Encode()
}

func printReference() {
	alt := fileName()
	alt = strings.TrimPrefix(alt, "img/")
	alt = strings.TrimSuffix(alt, fmt.Sprintf(".%s", *format))

	fmt.Printf(
		"|%s|%s|%s|%s|%d*%d|%s|![%s](%s)|\n",
		printArgs(),
		*theme,
		*style,
		*effect,
		size,
		size,
		*background,
		alt,
		fileName(),
	)
}

func printArgs() string {
	var args []string

	if *format != defaultFormat {
		args = append(args, fmt.Sprintf("-f=%s", *format))
	}
	if *theme != defaultTheme {
		args = append(args, fmt.Sprintf("-t=%s", *theme))
	}
	if *style != defaultStyle {
		args = append(args, fmt.Sprintf("-s=%s", *style))
	}
	if *effect != defaultEffect {
		args = append(args, fmt.Sprintf("-e=%s", *effect))
	}
	if *background != defaultBackground {
		args = append(args, fmt.Sprintf("-b=%s", *background))
	}
	if *delay != defaultDelay {
		args = append(args, fmt.Sprintf("-d=%d", *delay))
	}
	if *multiple != defaultMultiple {
		args = append(args, fmt.Sprintf("-m=%d", *multiple))
	}
	return strings.Join(args, " ")
}

func fileName() string {
	dir := "img"
	name := *theme

	if *style != portrait.BasicStyle {
		name = fmt.Sprintf("%s_%s", name, *style)
	}
	if *effect != "" {
		name = fmt.Sprintf("%s_%s", name, *effect)
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

			*theme = defaultTheme
			*style = defaultStyle
			*format = defaultFormat
			*effect = defaultEffect
			*background = defaultBackground
			*multiple = defaultMultiple
			*delay = defaultDelay

			for _, v := range strings.Split(ss[1], " ") {
				switch {
				case strings.HasPrefix(v, "-s="):
					*style = strings.TrimPrefix(v, "-s=")
				case strings.HasPrefix(v, "-t="):
					*theme = strings.TrimPrefix(v, "-t=")
				case strings.HasPrefix(v, "-f="):
					*format = strings.TrimPrefix(v, "-f=")
				case strings.HasPrefix(v, "-e="):
					*effect = strings.TrimPrefix(v, "-e=")
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

		switch {
		case fs.Text() == "## Args Example":
			afterArgsLine = true
		case strings.HasPrefix(fs.Text(), "|---"):
			afterHyphenLine = true
		}
	}

	return nil
}
