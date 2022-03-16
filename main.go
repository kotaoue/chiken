package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"

	"github.com/kotaoue/chiken/pkg/blueprint"
	"github.com/kotaoue/chiken/pkg/palette"
)

var (
	style      = flag.Int("s", 0, "style of rooster")
	multiple   = flag.Int("m", 1, "value to be multiplied by 32")
	background = flag.String("b", "", "background color. set with hex. example #ffffff. empty is transparent")
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
	c, err := hexToColor(*background)
	if err != nil {
		return err
	}
	fmt.Printf("size:%d multiple:%d background:%v\n", size, *multiple, c)

	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{size, size}})

	drawBG(img, c)
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
		name = fmt.Sprintf("%s_%d", name, *multiple)
	}
	if *background != "" {
		name = fmt.Sprintf("%s_%s", name, strings.ReplaceAll(*background, "#", ""))
	}
	return fmt.Sprintf("%s/%s.png", dir, name)
}

func drawBG(img *image.RGBA, col *color.RGBA) {
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

			if i != 0 {
				for my := 0; my < *multiple; my++ {
					for mx := 0; mx < *multiple; mx++ {
						img.Set(x**multiple+mx, y**multiple+my, cols[i])
					}
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
