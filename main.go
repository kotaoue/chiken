package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

var (
	size = flag.Int("s", 32, "icon size")
)

func init() {
	flag.Parse()
}

func main() {
	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Main() error {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{*size, *size}})
	drawBG(img, color.RGBA{255, 255, 255, 255})

	f, err := os.Create("ohyeah.png")
	if err != nil {
		return err
	}

	return png.Encode(f, img)
}

func drawBG(img *image.RGBA, col color.RGBA) {
	for x := 0; x < *size; x++ {
		for y := 0; y < *size; y++ {
			img.Set(x, y, col)
		}
	}
}
