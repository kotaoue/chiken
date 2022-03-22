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
	GreenTheme      = "green"
	MossGreenTheme  = "mossGreen"
	LightBlueTheme  = "lightBlue"
	BlueTheme       = "blue"
	BluePurpleTheme = "bluePurple"
	PurpleTheme     = "purple"
	PinkPurpleTheme = "pinkPurple"
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
	case GreenTheme:
		return t.green(), nil
	case MossGreenTheme:
		return t.mossGreen(), nil
	case LightBlueTheme:
		return t.lightBlue(), nil
	case BlueTheme:
		return t.blue(), nil
	case BluePurpleTheme:
		return t.bluePurple(), nil
	case PurpleTheme:
		return t.purple(), nil
	case PinkPurpleTheme:
		return t.pinkPurple(), nil
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

func (Theme) green() []color.Color {
	return []color.Color{
		color.RGBA{0, 0, 0, 0},         // 背景色
		color.RGBA{0, 0, 0, 255},       // 主線
		color.RGBA{41, 235, 68, 255},   // メインカラー
		color.RGBA{11, 158, 31, 255},   // メインカラー 影
		color.RGBA{41, 235, 68, 255},   // サブカラー
		color.RGBA{11, 158, 31, 255},   // サブカラー 影
		color.RGBA{255, 0, 0, 255},     // トサカ
		color.RGBA{255, 128, 128, 255}, // トサカ ハイライト
		color.RGBA{196, 0, 0, 255},     // トサカ 影
		color.RGBA{255, 255, 0, 255},   // くちばし
		color.RGBA{255, 255, 255, 255}, // くちばし ハイライト
		color.RGBA{255, 255, 0, 255},   // 足
	}
}

func (Theme) mossGreen() []color.Color {
	return []color.Color{
		color.RGBA{0, 0, 0, 0},         // 背景色
		color.RGBA{0, 0, 0, 255},       // 主線
		color.RGBA{21, 125, 38, 255},   // メインカラー
		color.RGBA{5, 69, 15, 255},     // メインカラー 影
		color.RGBA{21, 125, 38, 255},   // サブカラー
		color.RGBA{5, 69, 15, 255},     // サブカラー 影
		color.RGBA{255, 0, 0, 255},     // トサカ
		color.RGBA{255, 128, 128, 255}, // トサカ ハイライト
		color.RGBA{196, 0, 0, 255},     // トサカ 影
		color.RGBA{255, 255, 0, 255},   // くちばし
		color.RGBA{255, 255, 255, 255}, // くちばし ハイライト
		color.RGBA{255, 255, 0, 255},   // 足
	}
}

func (Theme) lightBlue() []color.Color {
	return []color.Color{
		color.RGBA{0, 0, 0, 0},         // 背景色
		color.RGBA{0, 0, 0, 255},       // 主線
		color.RGBA{86, 235, 207, 255},  // メインカラー
		color.RGBA{27, 158, 134, 255},  // メインカラー 影
		color.RGBA{86, 235, 207, 255},  // サブカラー
		color.RGBA{27, 158, 134, 255},  // サブカラー 影
		color.RGBA{255, 0, 0, 255},     // トサカ
		color.RGBA{255, 128, 128, 255}, // トサカ ハイライト
		color.RGBA{196, 0, 0, 255},     // トサカ 影
		color.RGBA{255, 255, 0, 255},   // くちばし
		color.RGBA{255, 255, 255, 255}, // くちばし ハイライト
		color.RGBA{255, 255, 0, 255},   // 足
	}
}

func (Theme) blue() []color.Color {
	return []color.Color{
		color.RGBA{0, 0, 0, 0},         // 背景色
		color.RGBA{0, 0, 0, 255},       // 主線
		color.RGBA{75, 166, 235, 255},  // メインカラー
		color.RGBA{19, 98, 158, 255},   // メインカラー 影
		color.RGBA{75, 166, 235, 255},  // サブカラー
		color.RGBA{19, 98, 158, 255},   // サブカラー 影
		color.RGBA{255, 0, 0, 255},     // トサカ
		color.RGBA{255, 128, 128, 255}, // トサカ ハイライト
		color.RGBA{196, 0, 0, 255},     // トサカ 影
		color.RGBA{255, 255, 0, 255},   // くちばし
		color.RGBA{255, 255, 255, 255}, // くちばし ハイライト
		color.RGBA{255, 255, 0, 255},   // 足
	}
}

func (Theme) bluePurple() []color.Color {
	return []color.Color{
		color.RGBA{0, 0, 0, 0},         // 背景色
		color.RGBA{0, 0, 0, 255},       // 主線
		color.RGBA{94, 108, 235, 255},  // メインカラー
		color.RGBA{31, 44, 158, 255},   // メインカラー 影
		color.RGBA{94, 108, 235, 255},  // サブカラー
		color.RGBA{31, 44, 158, 255},   // サブカラー 影
		color.RGBA{255, 0, 0, 255},     // トサカ
		color.RGBA{255, 128, 128, 255}, // トサカ ハイライト
		color.RGBA{196, 0, 0, 255},     // トサカ 影
		color.RGBA{255, 255, 0, 255},   // くちばし
		color.RGBA{255, 255, 255, 255}, // くちばし ハイライト
		color.RGBA{255, 255, 0, 255},   // 足
	}
}

func (Theme) purple() []color.Color {
	return []color.Color{
		color.RGBA{0, 0, 0, 0},         // 背景色
		color.RGBA{0, 0, 0, 255},       // 主線
		color.RGBA{145, 92, 235, 255},  // メインカラー
		color.RGBA{77, 30, 158, 255},   // メインカラー 影
		color.RGBA{145, 92, 235, 255},  // サブカラー
		color.RGBA{77, 30, 158, 255},   // サブカラー 影
		color.RGBA{255, 0, 0, 255},     // トサカ
		color.RGBA{255, 128, 128, 255}, // トサカ ハイライト
		color.RGBA{196, 0, 0, 255},     // トサカ 影
		color.RGBA{255, 255, 0, 255},   // くちばし
		color.RGBA{255, 255, 255, 255}, // くちばし ハイライト
		color.RGBA{255, 255, 0, 255},   // 足
	}
}

func (Theme) pinkPurple() []color.Color {
	return []color.Color{
		color.RGBA{0, 0, 0, 0},         // 背景色
		color.RGBA{0, 0, 0, 255},       // 主線
		color.RGBA{211, 91, 235, 255},  // メインカラー
		color.RGBA{136, 22, 158, 255},  // メインカラー 影
		color.RGBA{211, 91, 235, 255},  // サブカラー
		color.RGBA{136, 22, 158, 255},  // サブカラー 影
		color.RGBA{255, 0, 0, 255},     // トサカ
		color.RGBA{255, 128, 128, 255}, // トサカ ハイライト
		color.RGBA{196, 0, 0, 255},     // トサカ 影
		color.RGBA{255, 255, 0, 255},   // くちばし
		color.RGBA{255, 255, 255, 255}, // くちばし ハイライト
		color.RGBA{255, 255, 0, 255},   // 足
	}
}
