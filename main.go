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
	style    = flag.Int("s", 0, "style of rooster")
	multiple = flag.Int("m", 1, "value to be multiplied by 32")
	size     int
	baseSize int
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
	fmt.Printf("size:%d multiple:%d\n", size, *multiple)

	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{size, size}})

	drawBG(img, color.RGBA{255, 255, 255, 255})
	if err := drawImage(img); err != nil {
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
	plt := &palette.Palette{}
	name := plt.Name(*style)
	if *multiple > 1 {
		return fmt.Sprintf("%s/%s_%d.png", dir, name, *multiple)
	}
	return fmt.Sprintf("%s/%s.png", dir, name)
}

func drawBG(img *image.RGBA, col color.RGBA) {
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			img.Set(x, y, col)
		}
	}
}

func drawImage(img *image.RGBA) error {
	tpl, cols, err := fetchBlueprint()
	if err != nil {
		return err
	}

	for y := 0; y < baseSize; y++ {
		fmt.Print("{")
		for x := 0; x < baseSize; x++ {
			i := tpl[y][x]
			fmt.Print(i)

			for my := 0; my < *multiple; my++ {
				for mx := 0; mx < *multiple; mx++ {
					img.Set(x**multiple+mx, y**multiple+my, cols[i])
				}
			}
		}
		fmt.Println("}")
	}

	return nil
}

func fetchBlueprint() ([][]int, []color.RGBA, error) {
	bp := &blueprint.Blueprint{}
	p := &palette.Palette{}
	plt, err := p.Get(*style)
	if err != nil {
		return nil, nil, err
	}

	return bp.Get(blueprint.BasicStyle), plt, nil
}
