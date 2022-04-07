package portrait

import (
	"errors"
	"fmt"
	"image/color"
	"math"
	"strconv"
	"strings"
)

const (
	PartyTheme      = "party"
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
	PinkTheme       = "pink"
	RedTheme        = "red"
	OrangeTheme     = "orange"
	GrayTheme       = "gray"
	Player2Theme    = "player2"
	Player3Theme    = "player3"
	Player4Theme    = "player4"
	Player5Theme    = "player5"
)

type Theme struct{}

func (Theme) Adjust(theme string) (string, error) {
	if strings.HasPrefix(theme, PartyTheme) {
		ratio, err := strconv.Atoi(strings.TrimPrefix(theme, PartyTheme))
		if err != nil || ratio <= 0 {
			return "", fmt.Errorf("party expects a format of party{int}. but '%s' was specified", theme)
		}

		var ss []string
		angle := math.Round(360 / float64(ratio))
		for i := 0; i < 360; i += int(angle) {
			ss = append(ss, fmt.Sprintf("%s%d", PartyTheme, i))
		}
		s := strings.Join(ss, "-")
		return s, nil
	}

	return theme, nil
}

func (t Theme) Get(theme string) ([]color.Color, error) {
	if strings.HasPrefix(theme, PartyTheme) {
		return t.party(theme)
	}

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
	case PinkTheme:
		return t.pink(), nil
	case RedTheme:
		return t.red(), nil
	case OrangeTheme:
		return t.orange(), nil
	case GrayTheme:
		return t.gray(), nil
	case Player2Theme:
		return t.player2(), nil
	case Player3Theme:
		return t.player3(), nil
	case Player4Theme:
		return t.player4(), nil
	case Player5Theme:
		return t.player5(), nil
	}
	return nil, errors.New("not exist theme")
}

func (t Theme) party(s string) ([]color.Color, error) {
	i, err := strconv.Atoi(strings.TrimPrefix(s, PartyTheme))
	if err != nil || i < 0 {
		return nil, fmt.Errorf("party expects a format of party{int}. but '%s' was specified", s)
	}

	c := t.basic()

	rRad := float64(i) * math.Pi / float64(180)
	gRad := float64(i+120) * math.Pi / float64(180)
	bRad := float64(i+240) * math.Pi / float64(180)

	c[2] = color.RGBA{
		uint8(math.Abs(math.Round(255 * math.Sin(rRad)))),
		uint8(math.Abs(math.Round(255 * math.Sin(bRad)))),
		uint8(math.Abs(math.Round(255 * math.Sin(gRad)))),
		255,
	}
	c[3] = color.RGBA{
		uint8(math.Abs(math.Round(196 * math.Sin(rRad)))),
		uint8(math.Abs(math.Round(196 * math.Sin(bRad)))),
		uint8(math.Abs(math.Round(196 * math.Sin(gRad)))),
		255,
	}

	c[4] = c[2]
	c[5] = c[3]

	return c, nil
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

func (Theme) pink() []color.Color {
	return []color.Color{
		color.RGBA{0, 0, 0, 0},         // 背景色
		color.RGBA{0, 0, 0, 255},       // 主線
		color.RGBA{235, 91, 177, 255},  // メインカラー
		color.RGBA{158, 11, 98, 255},   // メインカラー 影
		color.RGBA{235, 91, 177, 255},  // サブカラー
		color.RGBA{158, 11, 98, 255},   // サブカラー 影
		color.RGBA{255, 0, 0, 255},     // トサカ
		color.RGBA{255, 128, 128, 255}, // トサカ ハイライト
		color.RGBA{196, 0, 0, 255},     // トサカ 影
		color.RGBA{255, 255, 0, 255},   // くちばし
		color.RGBA{255, 255, 255, 255}, // くちばし ハイライト
		color.RGBA{255, 255, 0, 255},   // 足
	}
}

func (Theme) red() []color.Color {
	return []color.Color{
		color.RGBA{0, 0, 0, 0},         // 背景色
		color.RGBA{0, 0, 0, 255},       // 主線
		color.RGBA{235, 33, 18, 255},   // メインカラー
		color.RGBA{158, 78, 73, 255},   // メインカラー 影
		color.RGBA{235, 33, 18, 255},   // サブカラー
		color.RGBA{158, 78, 73, 255},   // サブカラー 影
		color.RGBA{255, 0, 0, 255},     // トサカ
		color.RGBA{255, 128, 128, 255}, // トサカ ハイライト
		color.RGBA{196, 0, 0, 255},     // トサカ 影
		color.RGBA{255, 255, 0, 255},   // くちばし
		color.RGBA{255, 255, 255, 255}, // くちばし ハイライト
		color.RGBA{255, 255, 0, 255},   // 足
	}
}

func (Theme) orange() []color.Color {
	return []color.Color{
		color.RGBA{0, 0, 0, 0},         // 背景色
		color.RGBA{0, 0, 0, 255},       // 主線
		color.RGBA{255, 120, 29, 255},  // メインカラー
		color.RGBA{178, 103, 79, 255},  // メインカラー 影
		color.RGBA{255, 120, 29, 255},  // サブカラー
		color.RGBA{178, 103, 79, 255},  // サブカラー 影
		color.RGBA{255, 0, 0, 255},     // トサカ
		color.RGBA{255, 128, 128, 255}, // トサカ ハイライト
		color.RGBA{196, 0, 0, 255},     // トサカ 影
		color.RGBA{255, 255, 0, 255},   // くちばし
		color.RGBA{255, 255, 255, 255}, // くちばし ハイライト
		color.RGBA{255, 255, 0, 255},   // 足
	}
}

func (Theme) gray() []color.Color {
	return []color.Color{
		color.RGBA{0, 0, 0, 0},         // 背景色
		color.RGBA{0, 0, 0, 255},       // 主線
		color.RGBA{120, 120, 120, 255}, // メインカラー
		color.RGBA{79, 79, 79, 255},    // メインカラー 影
		color.RGBA{120, 120, 120, 255}, // サブカラー
		color.RGBA{79, 79, 79, 255},    // サブカラー 影
		color.RGBA{255, 0, 0, 255},     // トサカ
		color.RGBA{255, 128, 128, 255}, // トサカ ハイライト
		color.RGBA{196, 0, 0, 255},     // トサカ 影
		color.RGBA{255, 255, 0, 255},   // くちばし
		color.RGBA{255, 255, 255, 255}, // くちばし ハイライト
		color.RGBA{255, 255, 0, 255},   // 足
	}
}

func (Theme) player2() []color.Color {
	return []color.Color{
		color.RGBA{0, 0, 0, 0},         // 背景色
		color.RGBA{0, 0, 0, 255},       // 主線
		color.RGBA{255, 255, 255, 255}, // メインカラー
		color.RGBA{196, 196, 196, 255}, // メインカラー 影
		color.RGBA{255, 255, 255, 255}, // サブカラー
		color.RGBA{196, 196, 196, 255}, // サブカラー 影
		color.RGBA{0, 0, 255, 255},     // トサカ
		color.RGBA{64, 128, 255, 255},  // トサカ ハイライト
		color.RGBA{0, 0, 196, 255},     // トサカ 影
		color.RGBA{255, 255, 0, 255},   // くちばし
		color.RGBA{255, 255, 255, 255}, // くちばし ハイライト
		color.RGBA{255, 255, 0, 255},   // 足
	}
}

func (Theme) player3() []color.Color {
	return []color.Color{
		color.RGBA{0, 0, 0, 0},         // 背景色
		color.RGBA{0, 0, 0, 255},       // 主線
		color.RGBA{255, 255, 255, 255}, // メインカラー
		color.RGBA{196, 196, 196, 255}, // メインカラー 影
		color.RGBA{255, 255, 255, 255}, // サブカラー
		color.RGBA{196, 196, 196, 255}, // サブカラー 影
		color.RGBA{0, 196, 0, 255},     // トサカ
		color.RGBA{96, 196, 96, 255},   // トサカ ハイライト
		color.RGBA{0, 96, 0, 255},      // トサカ 影
		color.RGBA{255, 255, 0, 255},   // くちばし
		color.RGBA{255, 255, 255, 255}, // くちばし ハイライト
		color.RGBA{255, 255, 0, 255},   // 足
	}
}

func (Theme) player4() []color.Color {
	return []color.Color{
		color.RGBA{0, 0, 0, 0},         // 背景色
		color.RGBA{0, 0, 0, 255},       // 主線
		color.RGBA{255, 255, 255, 255}, // メインカラー
		color.RGBA{196, 196, 196, 255}, // メインカラー 影
		color.RGBA{255, 255, 255, 255}, // サブカラー
		color.RGBA{196, 196, 196, 255}, // サブカラー 影
		color.RGBA{255, 255, 0, 255},   // トサカ
		color.RGBA{128, 128, 0, 255},   // トサカ ハイライト
		color.RGBA{96, 96, 0, 255},     // トサカ 影
		color.RGBA{255, 255, 0, 255},   // くちばし
		color.RGBA{255, 255, 255, 255}, // くちばし ハイライト
		color.RGBA{255, 255, 0, 255},   // 足
	}
}

func (Theme) player5() []color.Color {
	return []color.Color{
		color.RGBA{0, 0, 0, 0},         // 背景色
		color.RGBA{0, 0, 0, 255},       // 主線
		color.RGBA{255, 255, 255, 255}, // メインカラー
		color.RGBA{196, 196, 196, 255}, // メインカラー 影
		color.RGBA{255, 255, 255, 255}, // サブカラー
		color.RGBA{196, 196, 196, 255}, // サブカラー 影
		color.RGBA{0, 255, 255, 255},   // トサカ
		color.RGBA{0, 128, 128, 255},   // トサカ ハイライト
		color.RGBA{96, 96, 0, 255},     // トサカ 影
		color.RGBA{255, 255, 0, 255},   // くちばし
		color.RGBA{255, 255, 255, 255}, // くちばし ハイライト
		color.RGBA{255, 255, 0, 255},   // 足
	}
}
