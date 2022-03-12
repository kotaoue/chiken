package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/kotaoue/chiken/pkg/blueprint"
)

var (
	size = flag.Int("s", 32, "size of output image")
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
	return Output()
}

func Output() error {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{*size, *size}})
	drawBG(img, color.RGBA{255, 255, 255, 255})

	drawImage(img)

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

func drawImage(img *image.RGBA) {
	bp := &blueprint.Blueprint{}
	tpl := bp.Get(blueprint.BasicStyle)

	cols := []color.RGBA{
		{0, 0, 0, 0},         // 背景色
		{0, 0, 0, 255},       // 主線
		{255, 255, 255, 255}, // メインカラー
		{128, 128, 128, 255}, // メインカラー 影
		{255, 0, 0, 255},     // トサカ
		{255, 128, 128, 255}, // トサカ ハイライト
		{128, 0, 0, 255},     // トサカ 影
		{0, 255, 255, 255},   // くちばし
		{255, 255, 255, 255}, // くちばし ハイライト
		{0, 255, 255, 255},   // 足
	}

	for x := 0; x < *size; x++ {
		fmt.Print("{")
		for y := 0; y < *size; y++ {
			i := tpl[x][y]
			fmt.Print(i)

			img.Set(x, y, cols[i])
		}
		fmt.Println("}")
	}
}
