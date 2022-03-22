package portrait

import (
	"errors"
	"image/color"
)

const (
	WhiteTheme      = "white"
	BrownTheme      = "brown"
	BlackTheme      = "black"
	BrownBlackTheme = "brownBlack"
	PandaTheme      = "panda"
	YellowTheme     = "yellow"
)

type Theme struct{}

func (t Theme) Get(theme string) ([]color.Color, error) {
	switch theme {
	case WhiteTheme:
		return t.basic(), nil
	case BrownTheme:
		return t.brown(), nil
	case BlackTheme:
		return t.black(), nil
	case BrownBlackTheme:
		return t.brownBlack(), nil
	case PandaTheme:
		return t.panda(), nil
	case YellowTheme:
		return t.yellow(), nil
	}
	return nil, errors.New("not exist theme")
}

func (Theme) basic() []color.Color {
	return []color.Color{
		color.RGBA{0, 0, 0, 0},         // 背景色
		color.RGBA{0, 0, 0, 255},       // 主線
		color.RGBA{255, 255, 255, 255}, // メインカラー
		color.RGBA{196, 196, 196, 255}, // メインカラー 影
		color.RGBA{255, 255, 255, 255}, // サブカラー
		color.RGBA{196, 196, 196, 255}, // サブカラー 影
		color.RGBA{255, 0, 0, 255},     // トサカ
		color.RGBA{255, 128, 128, 255}, // トサカ ハイライト
		color.RGBA{196, 0, 0, 255},     // トサカ 影
		color.RGBA{255, 255, 0, 255},   // くちばし
		color.RGBA{255, 255, 255, 255}, // くちばし ハイライト
		color.RGBA{255, 255, 0, 255},   // 足
	}
}

func (Theme) brown() []color.Color {
	return []color.Color{
		color.RGBA{0, 0, 0, 0},         // 背景色
		color.RGBA{0, 0, 0, 255},       // 主線
		color.RGBA{204, 65, 37, 255},   // メインカラー
		color.RGBA{166, 28, 0, 255},    // メインカラー 影
		color.RGBA{204, 65, 37, 255},   // サブカラー
		color.RGBA{166, 28, 0, 255},    // サブカラー 影
		color.RGBA{255, 0, 0, 255},     // トサカ
		color.RGBA{255, 128, 128, 255}, // トサカ ハイライト
		color.RGBA{128, 0, 0, 255},     // トサカ 影
		color.RGBA{255, 255, 0, 255},   // くちばし
		color.RGBA{255, 255, 255, 255}, // くちばし ハイライト
		color.RGBA{255, 255, 0, 255},   // 足
	}
}

func (Theme) brownBlack() []color.Color {
	return []color.Color{
		color.RGBA{0, 0, 0, 0},         // 背景色
		color.RGBA{0, 0, 0, 255},       // 主線
		color.RGBA{204, 65, 37, 255},   // メインカラー
		color.RGBA{166, 28, 0, 255},    // メインカラー 影
		color.RGBA{91, 15, 0, 255},     // サブカラー
		color.RGBA{32, 8, 0, 255},      // サブカラー 影
		color.RGBA{255, 0, 0, 255},     // トサカ
		color.RGBA{255, 128, 128, 255}, // トサカ ハイライト
		color.RGBA{128, 0, 0, 255},     // トサカ 影
		color.RGBA{255, 255, 0, 255},   // くちばし
		color.RGBA{255, 255, 255, 255}, // くちばし ハイライト
		color.RGBA{255, 255, 0, 255},   // 足
	}
}

func (Theme) black() []color.Color {
	return []color.Color{
		color.RGBA{0, 0, 0, 0},         // 背景色
		color.RGBA{0, 0, 0, 255},       // 主線
		color.RGBA{64, 64, 64, 255},    // メインカラー
		color.RGBA{48, 48, 48, 255},    // メインカラー 影
		color.RGBA{64, 64, 64, 255},    // サブカラー
		color.RGBA{48, 48, 48, 255},    // サブカラー 影
		color.RGBA{255, 0, 0, 255},     // トサカ
		color.RGBA{255, 128, 128, 255}, // トサカ ハイライト
		color.RGBA{128, 0, 0, 255},     // トサカ 影
		color.RGBA{255, 255, 0, 255},   // くちばし
		color.RGBA{255, 255, 255, 255}, // くちばし ハイライト
		color.RGBA{255, 255, 0, 255},   // 足
	}
}

func (Theme) panda() []color.Color {
	return []color.Color{
		color.RGBA{0, 0, 0, 0},         // 背景色
		color.RGBA{0, 0, 0, 255},       // 主線
		color.RGBA{255, 255, 255, 255}, // メインカラー
		color.RGBA{196, 196, 196, 255}, // メインカラー 影
		color.RGBA{64, 64, 64, 255},    // サブカラー
		color.RGBA{48, 48, 48, 255},    // サブカラー 影
		color.RGBA{255, 0, 0, 255},     // トサカ
		color.RGBA{255, 128, 128, 255}, // トサカ ハイライト
		color.RGBA{196, 0, 0, 255},     // トサカ 影
		color.RGBA{255, 255, 0, 255},   // くちばし
		color.RGBA{255, 255, 255, 255}, // くちばし ハイライト
		color.RGBA{255, 255, 0, 255},   // 足
	}
}

func (Theme) yellow() []color.Color {
	return []color.Color{
		color.RGBA{0, 0, 0, 0},         // 背景色
		color.RGBA{0, 0, 0, 255},       // 主線
		color.RGBA{235, 220, 49, 255},  // メインカラー
		color.RGBA{158, 146, 17, 255},  // メインカラー 影
		color.RGBA{235, 220, 49, 255},  // サブカラー
		color.RGBA{158, 146, 17, 255},  // サブカラー 影
		color.RGBA{255, 0, 0, 255},     // トサカ
		color.RGBA{255, 128, 128, 255}, // トサカ ハイライト
		color.RGBA{196, 0, 0, 255},     // トサカ 影
		color.RGBA{255, 255, 0, 255},   // くちばし
		color.RGBA{255, 255, 255, 255}, // くちばし ハイライト
		color.RGBA{255, 255, 0, 255},   // 足
	}
}
