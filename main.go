package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/kotaoue/chiken/pkg/blueprint"
	"github.com/kotaoue/chiken/pkg/palette"
)

var (
	size = flag.Int("s", 32, "size of output image. request multiples of 32")
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
	return output()
}

func output() error {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{*size, *size}})

	drawBG(img, color.RGBA{255, 255, 255, 255})
	drawImage(img)

	f, err := os.Create("img/basic.png")
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

func drawImage(img *image.RGBA) {
	tpl, cols := fetchBlueprint()

	for y := 0; y < *size; y++ {
		fmt.Print("{")
		for x := 0; x < *size; x++ {
			i := tpl[y][x]
			fmt.Print(i)

			img.Set(x, y, cols[i])
		}
		fmt.Println("}")
	}
}

func fetchBlueprint() ([][]int, []color.RGBA) {
	bp := &blueprint.Blueprint{}
	plt := &palette.Palette{}

	return bp.Get(blueprint.BasicStyle), plt.Get(palette.BasicStyle)
}
