package palette

import (
	"errors"
	"image/color"
)

const (
	WhiteTheme = "white"
	BlackTheme = "black"
)

func Get(theme string) ([]color.RGBA, error) {
	switch theme {
	case WhiteTheme:
		return basic(), nil
	case BlackTheme:
		return black(), nil
	}
	return nil, errors.New("not exist palette")
}

func basic() []color.RGBA {
	return []color.RGBA{
		{0, 0, 0, 0},         // 背景色
		{0, 0, 0, 255},       // 主線
		{255, 255, 255, 255}, // メインカラー
		{128, 128, 128, 255}, // メインカラー 影
		{255, 0, 0, 255},     // トサカ
		{255, 128, 128, 255}, // トサカ ハイライト
		{128, 0, 0, 255},     // トサカ 影
		{255, 255, 0, 255},   // くちばし
		{255, 255, 255, 255}, // くちばし ハイライト
		{255, 255, 0, 255},   // 足
	}
}

func black() []color.RGBA {
	return []color.RGBA{
		{0, 0, 0, 0},         // 背景色
		{0, 0, 0, 255},       // 主線
		{64, 64, 64, 255},    // メインカラー
		{32, 32, 32, 255},    // メインカラー 影
		{255, 0, 0, 255},     // トサカ
		{255, 128, 128, 255}, // トサカ ハイライト
		{128, 0, 0, 255},     // トサカ 影
		{255, 255, 0, 255},   // くちばし
		{255, 255, 255, 255}, // くちばし ハイライト
		{255, 255, 0, 255},   // 足
	}
}
