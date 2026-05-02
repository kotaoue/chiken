package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/kotaoue/chiken/pkg/portrait"
)

const (
	defaultTheme            = portrait.WhiteTheme
	defaultStyle            = portrait.BasicStyle
	defaultFormat           = "png"
	defaultEffect           = ""
	defaultBackground       = "transparent"
	defaultName             = ""
	defaultMultiple         = 1
	defaultDelay            = 0
	defaultText             = ""
	defaultTextColor        = ""
	defaultTextFontSize     = 0
	defaultTextFont         = "regular"
	defaultBalloon          = false
	defaultBalloonLineColor = "#000000"
	defaultBalloonBgColor   = "#ffffff"
)

var (
	theme            string
	style            string
	format           string
	effect           string
	background       string
	name             string
	text             string
	textColor        string
	textFontSize     int
	textFont         string
	multiple         int
	delay            int
	verbose          bool
	dump             bool
	balloon          bool
	balloonLineColor string
	balloonBgColor   string
	size             int
	baseSize         = 32
	out              io.Writer
)

func init() {
	rootCmd.Flags().StringVarP(&theme, "theme", "t", defaultTheme, "theme color of rooster")
	rootCmd.Flags().StringVarP(&style, "style", "s", defaultStyle, "style of rooster")
	rootCmd.Flags().StringVarP(&format, "format", "f", defaultFormat, "format of output image")
	rootCmd.Flags().StringVarP(&effect, "effect", "e", defaultEffect, "set visual effects")
	rootCmd.Flags().StringVarP(&background, "background", "b", defaultBackground, "background color. set with hex. example #ffffff. empty is transparent")
	rootCmd.Flags().StringVarP(&name, "name", "n", defaultName, "name of output image")
	rootCmd.Flags().StringVarP(&text, "text", "T", defaultText, "text to display alongside the image")
	rootCmd.Flags().StringVarP(&textColor, "text-color", "c", defaultTextColor, "text color in hex format. example #ff0000")
	rootCmd.Flags().IntVar(&textFontSize, "text-font-size", defaultTextFontSize, "font size for text rendering. 0 uses the default 7x13 bitmap font")
	rootCmd.Flags().StringVar(&textFont, "text-font", defaultTextFont, fmt.Sprintf("font for text rendering. available: %s", strings.Join(portrait.ListFonts(), ", ")))
	rootCmd.Flags().IntVarP(&multiple, "multiple", "m", defaultMultiple, "value to be multiplied by 32")
	rootCmd.Flags().IntVarP(&delay, "delay", "d", defaultDelay, "delay time for gif")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "printing verbose output")
	rootCmd.Flags().BoolVar(&dump, "dump", false, "re encode from Args Example on README")
	rootCmd.Flags().BoolVarP(&balloon, "balloon", "B", defaultBalloon, "display text in a speech balloon with 8-bit style")
	rootCmd.Flags().StringVar(&balloonLineColor, "balloon-line-color", defaultBalloonLineColor, "balloon border color in hex format. example #000000")
	rootCmd.Flags().StringVar(&balloonBgColor, "balloon-bg-color", defaultBalloonBgColor, "balloon background color in hex format. example #ffffff")
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
