package main

import (
	"bufio"
	"errors"
	"fmt"
	"image/color"
	"os"
	"strconv"
	"strings"

	"github.com/kotaoue/chiken/pkg/cutil"
	"github.com/kotaoue/chiken/pkg/portrait"
	"github.com/spf13/cobra"
)

const (
	defaultTheme      = portrait.WhiteTheme
	defaultStyle      = portrait.BasicStyle
	defaultFormat     = "png"
	defaultEffect     = ""
	defaultBackground = "transparent"
	defaultName       = ""
	defaultMultiple   = 1
	defaultDelay      = 0
)

var (
	theme      string
	style      string
	format     string
	effect     string
	background string
	name       string
	multiple   int
	delay      int
	verbose    bool
	dump       bool
	size       int
	baseSize   = 32
)

var rootCmd = &cobra.Command{
	Use:   "chiken",
	Short: "A rooster image generator",
	Long:  `A CLI tool for generating rooster images with various themes, styles, and effects.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if dump {
			return reOutputs()
		}
		return output()
	},
}

func init() {
	rootCmd.Flags().StringVarP(&theme, "theme", "t", defaultTheme, "theme color of rooster")
	rootCmd.Flags().StringVarP(&style, "style", "s", defaultStyle, "style of rooster")
	rootCmd.Flags().StringVarP(&format, "format", "f", defaultFormat, "format of output image")
	rootCmd.Flags().StringVarP(&effect, "effect", "e", defaultEffect, "set visual effects")
	rootCmd.Flags().StringVarP(&background, "background", "b", defaultBackground, "background color. set with hex. example #ffffff. empty is transparent")
	rootCmd.Flags().StringVarP(&name, "name", "n", defaultName, "name of output image")
	rootCmd.Flags().IntVarP(&multiple, "multiple", "m", defaultMultiple, "value to be multiplied by 32")
	rootCmd.Flags().IntVarP(&delay, "delay", "d", defaultDelay, "delay time for gif")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "printing verbose output")
	rootCmd.Flags().BoolVar(&dump, "dump", false, "re encode from Args Example on README")
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func output() error {
	size = baseSize * multiple

	if err := checkFormat(format); err != nil {
		return err
	}

	c, err := cutil.HexToColor(background)
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
			Multiple:        multiple,
			Style:           style,
			Theme:           theme,
			BackgroundColor: c,
			Format:          format,
			Effect:          effect,
			Delay:           delay,
			FileName:        fileName(),
			Verbose:         verbose,
		},
	)

	return p.Encode()
}

func printReference() {
	alt := fileName()
	alt = strings.TrimPrefix(alt, "img/")
	alt = strings.TrimSuffix(alt, fmt.Sprintf(".%s", format))

	fmt.Printf(
		"|%s|%s|%s|%s|%d*%d|%s|![%s](%s)|\n",
		printArgs(),
		theme,
		style,
		effect,
		size,
		size,
		background,
		alt,
		fileName(),
	)
}

func printArgs() string {
	var args []string

	if format != defaultFormat {
		args = append(args, fmt.Sprintf("-f=%s", format))
	}
	if theme != defaultTheme {
		args = append(args, fmt.Sprintf("-t=%s", theme))
	}
	if style != defaultStyle {
		args = append(args, fmt.Sprintf("-s=%s", style))
	}
	if effect != defaultEffect {
		args = append(args, fmt.Sprintf("-e=%s", effect))
	}
	if background != defaultBackground {
		args = append(args, fmt.Sprintf("-b=%s", background))
	}
	if delay != defaultDelay {
		args = append(args, fmt.Sprintf("-d=%d", delay))
	}
	if multiple != defaultMultiple {
		args = append(args, fmt.Sprintf("-m=%d", multiple))
	}
	if name != defaultName {
		args = append(args, fmt.Sprintf("-n=%s", name))
	}
	return strings.Join(args, " ")
}

func fileName() string {
	dir := "img"
	if name != "" {
		return fmt.Sprintf("%s/%s.%s", dir, name, format)
	}

	fileName := theme

	if style != defaultStyle {
		fileName = fmt.Sprintf("%s_%s", fileName, style)
	}
	if effect != defaultEffect {
		fileName = fmt.Sprintf("%s_%s", fileName, effect)
	}
	if multiple != defaultMultiple {
		fileName = fmt.Sprintf("%s_%d", fileName, multiple)
	}
	if background != defaultBackground {
		fileName = fmt.Sprintf("%s_%s", fileName, strings.ReplaceAll(background, "#", ""))
	}
	if delay != defaultDelay {
		fileName = fmt.Sprintf("%s_delay%d", fileName, delay)
	}
	return fmt.Sprintf("%s/%s.%s", dir, fileName, format)
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

			theme = defaultTheme
			style = defaultStyle
			format = defaultFormat
			effect = defaultEffect
			background = defaultBackground
			name = defaultName
			multiple = defaultMultiple
			delay = defaultDelay

			for _, v := range strings.Split(ss[1], " ") {
				switch {
				case strings.HasPrefix(v, "-s="):
					style = strings.TrimPrefix(v, "-s=")
				case strings.HasPrefix(v, "-t="):
					theme = strings.TrimPrefix(v, "-t=")
				case strings.HasPrefix(v, "-f="):
					format = strings.TrimPrefix(v, "-f=")
				case strings.HasPrefix(v, "-e="):
					effect = strings.TrimPrefix(v, "-e=")
				case strings.HasPrefix(v, "-b="):
					background = strings.TrimPrefix(v, "-b=")
				case strings.HasPrefix(v, "-n="):
					name = strings.TrimPrefix(v, "-n=")
				case strings.HasPrefix(v, "-m="):
					i, err := strconv.Atoi(strings.TrimPrefix(v, "-m="))
					if err != nil {
						return err
					}
					multiple = i
				case strings.HasPrefix(v, "-d="):
					i, err := strconv.Atoi(strings.TrimPrefix(v, "-d="))
					if err != nil {
						return err
					}
					delay = i
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
