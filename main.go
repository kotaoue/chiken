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

	tc, err := cutil.HexToColor(textColor)
	if err != nil {
		return err
	}

	if err := encode(c, tc); err != nil {
		return err
	}

	printReference()

	return nil
}

func encode(c *color.RGBA, tc *color.RGBA) error {
	blc, err := cutil.HexToColor(balloonLineColor)
	if err != nil {
		return err
	}
	bbc, err := cutil.HexToColor(balloonBgColor)
	if err != nil {
		return err
	}

	p := portrait.NewPortrait(
		portrait.Options{
			Size:               size,
			BaseSize:           baseSize,
			Multiple:           multiple,
			Style:              style,
			Theme:              theme,
			BackgroundColor:    c,
			Format:             format,
			Effect:             effect,
			Delay:              delay,
			FileName:           fileName(),
			Verbose:            verbose,
			Output:             out,
			Text:               text,
			TextColor:          tc,
			TextFontSize:       textFontSize,
			TextFont:           textFont,
			Balloon:            balloon,
			BalloonBorderColor: blc,
			BalloonFillColor:   bbc,
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
	if text != defaultText {
		args = append(args, fmt.Sprintf("-T=%s", text))
	}
	if textColor != defaultTextColor {
		args = append(args, fmt.Sprintf("-c=%s", textColor))
	}
	if textFontSize != defaultTextFontSize {
		args = append(args, fmt.Sprintf("--text-font-size=%d", textFontSize))
	}
	if textFont != defaultTextFont {
		args = append(args, fmt.Sprintf("--text-font=%s", textFont))
	}
	if balloon != defaultBalloon {
		args = append(args, "--balloon")
	}
	if balloonLineColor != defaultBalloonLineColor {
		args = append(args, fmt.Sprintf("--balloon-line-color=%s", balloonLineColor))
	}
	if balloonBgColor != defaultBalloonBgColor {
		args = append(args, fmt.Sprintf("--balloon-bg-color=%s", balloonBgColor))
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
			if !strings.HasPrefix(fs.Text(), "|") {
				afterHyphenLine = false
				continue
			}

			ss := strings.Split(fs.Text(), "|")
			if len(ss) < 3 {
				continue
			}

			theme = defaultTheme
			style = defaultStyle
			format = defaultFormat
			effect = defaultEffect
			background = defaultBackground
			name = defaultName
			text = defaultText
			textColor = defaultTextColor
			textFontSize = defaultTextFontSize
			textFont = defaultTextFont
			multiple = defaultMultiple
			delay = defaultDelay
			balloon = defaultBalloon
			balloonLineColor = defaultBalloonLineColor
			balloonBgColor = defaultBalloonBgColor

			for _, v := range strings.Split(ss[2], " ") {
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
				case strings.HasPrefix(v, "-T="):
					text = strings.TrimPrefix(v, "-T=")
				case strings.HasPrefix(v, "-c="):
					textColor = strings.TrimPrefix(v, "-c=")
				case strings.HasPrefix(v, "--text-font-size="):
					i, err := strconv.Atoi(strings.TrimPrefix(v, "--text-font-size="))
					if err != nil {
						return err
					}
					textFontSize = i
				case strings.HasPrefix(v, "--text-font="):
					textFont = strings.TrimPrefix(v, "--text-font=")
				case v == "--balloon":
					balloon = true
				case strings.HasPrefix(v, "--balloon-line-color="):
					balloonLineColor = strings.TrimPrefix(v, "--balloon-line-color=")
				case strings.HasPrefix(v, "--balloon-bg-color="):
					balloonBgColor = strings.TrimPrefix(v, "--balloon-bg-color=")
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
