package gawk

// This file was generated.  Don't edit it directly.

import (
	"github.com/jwalton/gawk/pkg/ansistyles"
)

// Blue returns a string where the color is blue.
func Blue(str ...string) string {
	return rootBuilder.Blue(str...)
}

// WithBlue returns a Builder that generates strings where the color is blue,
// and further styles can be applied via chaining.
func WithBlue() *Builder {
	return rootBuilder.WithBlue()
}

// Blue returns a string where the color is blue, in addition to other styles from this builder.
func (builder *Builder) Blue(str ...string) string {
	return builder.WithBlue().applyStyle(str...)
}

// WithBlue returns a Builder that generates strings where the color is blue,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithBlue() *Builder {
	if builder.blue == nil {
		builder.blue = createBuilder(builder, ansistyles.Blue.Open, ansistyles.Blue.Close)
	}
	return builder.blue
}

// BlackBright returns a string where the color is blackBright.
func BlackBright(str ...string) string {
	return rootBuilder.BlackBright(str...)
}

// WithBlackBright returns a Builder that generates strings where the color is blackBright,
// and further styles can be applied via chaining.
func WithBlackBright() *Builder {
	return rootBuilder.WithBlackBright()
}

// BlackBright returns a string where the color is blackBright, in addition to other styles from this builder.
func (builder *Builder) BlackBright(str ...string) string {
	return builder.WithBlackBright().applyStyle(str...)
}

// WithBlackBright returns a Builder that generates strings where the color is blackBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithBlackBright() *Builder {
	if builder.blackBright == nil {
		builder.blackBright = createBuilder(builder, ansistyles.BlackBright.Open, ansistyles.BlackBright.Close)
	}
	return builder.blackBright
}

// MagentaBright returns a string where the color is magentaBright.
func MagentaBright(str ...string) string {
	return rootBuilder.MagentaBright(str...)
}

// WithMagentaBright returns a Builder that generates strings where the color is magentaBright,
// and further styles can be applied via chaining.
func WithMagentaBright() *Builder {
	return rootBuilder.WithMagentaBright()
}

// MagentaBright returns a string where the color is magentaBright, in addition to other styles from this builder.
func (builder *Builder) MagentaBright(str ...string) string {
	return builder.WithMagentaBright().applyStyle(str...)
}

// WithMagentaBright returns a Builder that generates strings where the color is magentaBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithMagentaBright() *Builder {
	if builder.magentaBright == nil {
		builder.magentaBright = createBuilder(builder, ansistyles.MagentaBright.Open, ansistyles.MagentaBright.Close)
	}
	return builder.magentaBright
}

// CyanBright returns a string where the color is cyanBright.
func CyanBright(str ...string) string {
	return rootBuilder.CyanBright(str...)
}

// WithCyanBright returns a Builder that generates strings where the color is cyanBright,
// and further styles can be applied via chaining.
func WithCyanBright() *Builder {
	return rootBuilder.WithCyanBright()
}

// CyanBright returns a string where the color is cyanBright, in addition to other styles from this builder.
func (builder *Builder) CyanBright(str ...string) string {
	return builder.WithCyanBright().applyStyle(str...)
}

// WithCyanBright returns a Builder that generates strings where the color is cyanBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithCyanBright() *Builder {
	if builder.cyanBright == nil {
		builder.cyanBright = createBuilder(builder, ansistyles.CyanBright.Open, ansistyles.CyanBright.Close)
	}
	return builder.cyanBright
}

// Grey returns a string where the color is grey.
func Grey(str ...string) string {
	return rootBuilder.Grey(str...)
}

// WithGrey returns a Builder that generates strings where the color is grey,
// and further styles can be applied via chaining.
func WithGrey() *Builder {
	return rootBuilder.WithGrey()
}

// Grey returns a string where the color is grey, in addition to other styles from this builder.
func (builder *Builder) Grey(str ...string) string {
	return builder.WithGrey().applyStyle(str...)
}

// WithGrey returns a Builder that generates strings where the color is grey,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithGrey() *Builder {
	if builder.grey == nil {
		builder.grey = createBuilder(builder, ansistyles.Grey.Open, ansistyles.Grey.Close)
	}
	return builder.grey
}

// Black returns a string where the color is black.
func Black(str ...string) string {
	return rootBuilder.Black(str...)
}

// WithBlack returns a Builder that generates strings where the color is black,
// and further styles can be applied via chaining.
func WithBlack() *Builder {
	return rootBuilder.WithBlack()
}

// Black returns a string where the color is black, in addition to other styles from this builder.
func (builder *Builder) Black(str ...string) string {
	return builder.WithBlack().applyStyle(str...)
}

// WithBlack returns a Builder that generates strings where the color is black,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithBlack() *Builder {
	if builder.black == nil {
		builder.black = createBuilder(builder, ansistyles.Black.Open, ansistyles.Black.Close)
	}
	return builder.black
}

// Red returns a string where the color is red.
func Red(str ...string) string {
	return rootBuilder.Red(str...)
}

// WithRed returns a Builder that generates strings where the color is red,
// and further styles can be applied via chaining.
func WithRed() *Builder {
	return rootBuilder.WithRed()
}

// Red returns a string where the color is red, in addition to other styles from this builder.
func (builder *Builder) Red(str ...string) string {
	return builder.WithRed().applyStyle(str...)
}

// WithRed returns a Builder that generates strings where the color is red,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithRed() *Builder {
	if builder.red == nil {
		builder.red = createBuilder(builder, ansistyles.Red.Open, ansistyles.Red.Close)
	}
	return builder.red
}

// Green returns a string where the color is green.
func Green(str ...string) string {
	return rootBuilder.Green(str...)
}

// WithGreen returns a Builder that generates strings where the color is green,
// and further styles can be applied via chaining.
func WithGreen() *Builder {
	return rootBuilder.WithGreen()
}

// Green returns a string where the color is green, in addition to other styles from this builder.
func (builder *Builder) Green(str ...string) string {
	return builder.WithGreen().applyStyle(str...)
}

// WithGreen returns a Builder that generates strings where the color is green,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithGreen() *Builder {
	if builder.green == nil {
		builder.green = createBuilder(builder, ansistyles.Green.Open, ansistyles.Green.Close)
	}
	return builder.green
}

// Yellow returns a string where the color is yellow.
func Yellow(str ...string) string {
	return rootBuilder.Yellow(str...)
}

// WithYellow returns a Builder that generates strings where the color is yellow,
// and further styles can be applied via chaining.
func WithYellow() *Builder {
	return rootBuilder.WithYellow()
}

// Yellow returns a string where the color is yellow, in addition to other styles from this builder.
func (builder *Builder) Yellow(str ...string) string {
	return builder.WithYellow().applyStyle(str...)
}

// WithYellow returns a Builder that generates strings where the color is yellow,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithYellow() *Builder {
	if builder.yellow == nil {
		builder.yellow = createBuilder(builder, ansistyles.Yellow.Open, ansistyles.Yellow.Close)
	}
	return builder.yellow
}

// Magenta returns a string where the color is magenta.
func Magenta(str ...string) string {
	return rootBuilder.Magenta(str...)
}

// WithMagenta returns a Builder that generates strings where the color is magenta,
// and further styles can be applied via chaining.
func WithMagenta() *Builder {
	return rootBuilder.WithMagenta()
}

// Magenta returns a string where the color is magenta, in addition to other styles from this builder.
func (builder *Builder) Magenta(str ...string) string {
	return builder.WithMagenta().applyStyle(str...)
}

// WithMagenta returns a Builder that generates strings where the color is magenta,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithMagenta() *Builder {
	if builder.magenta == nil {
		builder.magenta = createBuilder(builder, ansistyles.Magenta.Open, ansistyles.Magenta.Close)
	}
	return builder.magenta
}

// BlueBright returns a string where the color is blueBright.
func BlueBright(str ...string) string {
	return rootBuilder.BlueBright(str...)
}

// WithBlueBright returns a Builder that generates strings where the color is blueBright,
// and further styles can be applied via chaining.
func WithBlueBright() *Builder {
	return rootBuilder.WithBlueBright()
}

// BlueBright returns a string where the color is blueBright, in addition to other styles from this builder.
func (builder *Builder) BlueBright(str ...string) string {
	return builder.WithBlueBright().applyStyle(str...)
}

// WithBlueBright returns a Builder that generates strings where the color is blueBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithBlueBright() *Builder {
	if builder.blueBright == nil {
		builder.blueBright = createBuilder(builder, ansistyles.BlueBright.Open, ansistyles.BlueBright.Close)
	}
	return builder.blueBright
}

// Cyan returns a string where the color is cyan.
func Cyan(str ...string) string {
	return rootBuilder.Cyan(str...)
}

// WithCyan returns a Builder that generates strings where the color is cyan,
// and further styles can be applied via chaining.
func WithCyan() *Builder {
	return rootBuilder.WithCyan()
}

// Cyan returns a string where the color is cyan, in addition to other styles from this builder.
func (builder *Builder) Cyan(str ...string) string {
	return builder.WithCyan().applyStyle(str...)
}

// WithCyan returns a Builder that generates strings where the color is cyan,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithCyan() *Builder {
	if builder.cyan == nil {
		builder.cyan = createBuilder(builder, ansistyles.Cyan.Open, ansistyles.Cyan.Close)
	}
	return builder.cyan
}

// White returns a string where the color is white.
func White(str ...string) string {
	return rootBuilder.White(str...)
}

// WithWhite returns a Builder that generates strings where the color is white,
// and further styles can be applied via chaining.
func WithWhite() *Builder {
	return rootBuilder.WithWhite()
}

// White returns a string where the color is white, in addition to other styles from this builder.
func (builder *Builder) White(str ...string) string {
	return builder.WithWhite().applyStyle(str...)
}

// WithWhite returns a Builder that generates strings where the color is white,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithWhite() *Builder {
	if builder.white == nil {
		builder.white = createBuilder(builder, ansistyles.White.Open, ansistyles.White.Close)
	}
	return builder.white
}

// RedBright returns a string where the color is redBright.
func RedBright(str ...string) string {
	return rootBuilder.RedBright(str...)
}

// WithRedBright returns a Builder that generates strings where the color is redBright,
// and further styles can be applied via chaining.
func WithRedBright() *Builder {
	return rootBuilder.WithRedBright()
}

// RedBright returns a string where the color is redBright, in addition to other styles from this builder.
func (builder *Builder) RedBright(str ...string) string {
	return builder.WithRedBright().applyStyle(str...)
}

// WithRedBright returns a Builder that generates strings where the color is redBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithRedBright() *Builder {
	if builder.redBright == nil {
		builder.redBright = createBuilder(builder, ansistyles.RedBright.Open, ansistyles.RedBright.Close)
	}
	return builder.redBright
}

// WhiteBright returns a string where the color is whiteBright.
func WhiteBright(str ...string) string {
	return rootBuilder.WhiteBright(str...)
}

// WithWhiteBright returns a Builder that generates strings where the color is whiteBright,
// and further styles can be applied via chaining.
func WithWhiteBright() *Builder {
	return rootBuilder.WithWhiteBright()
}

// WhiteBright returns a string where the color is whiteBright, in addition to other styles from this builder.
func (builder *Builder) WhiteBright(str ...string) string {
	return builder.WithWhiteBright().applyStyle(str...)
}

// WithWhiteBright returns a Builder that generates strings where the color is whiteBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithWhiteBright() *Builder {
	if builder.whiteBright == nil {
		builder.whiteBright = createBuilder(builder, ansistyles.WhiteBright.Open, ansistyles.WhiteBright.Close)
	}
	return builder.whiteBright
}

// GreenBright returns a string where the color is greenBright.
func GreenBright(str ...string) string {
	return rootBuilder.GreenBright(str...)
}

// WithGreenBright returns a Builder that generates strings where the color is greenBright,
// and further styles can be applied via chaining.
func WithGreenBright() *Builder {
	return rootBuilder.WithGreenBright()
}

// GreenBright returns a string where the color is greenBright, in addition to other styles from this builder.
func (builder *Builder) GreenBright(str ...string) string {
	return builder.WithGreenBright().applyStyle(str...)
}

// WithGreenBright returns a Builder that generates strings where the color is greenBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithGreenBright() *Builder {
	if builder.greenBright == nil {
		builder.greenBright = createBuilder(builder, ansistyles.GreenBright.Open, ansistyles.GreenBright.Close)
	}
	return builder.greenBright
}

// YellowBright returns a string where the color is yellowBright.
func YellowBright(str ...string) string {
	return rootBuilder.YellowBright(str...)
}

// WithYellowBright returns a Builder that generates strings where the color is yellowBright,
// and further styles can be applied via chaining.
func WithYellowBright() *Builder {
	return rootBuilder.WithYellowBright()
}

// YellowBright returns a string where the color is yellowBright, in addition to other styles from this builder.
func (builder *Builder) YellowBright(str ...string) string {
	return builder.WithYellowBright().applyStyle(str...)
}

// WithYellowBright returns a Builder that generates strings where the color is yellowBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithYellowBright() *Builder {
	if builder.yellowBright == nil {
		builder.yellowBright = createBuilder(builder, ansistyles.YellowBright.Open, ansistyles.YellowBright.Close)
	}
	return builder.yellowBright
}

// Gray returns a string where the color is gray.
func Gray(str ...string) string {
	return rootBuilder.Gray(str...)
}

// WithGray returns a Builder that generates strings where the color is gray,
// and further styles can be applied via chaining.
func WithGray() *Builder {
	return rootBuilder.WithGray()
}

// Gray returns a string where the color is gray, in addition to other styles from this builder.
func (builder *Builder) Gray(str ...string) string {
	return builder.WithGray().applyStyle(str...)
}

// WithGray returns a Builder that generates strings where the color is gray,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithGray() *Builder {
	if builder.gray == nil {
		builder.gray = createBuilder(builder, ansistyles.Gray.Open, ansistyles.Gray.Close)
	}
	return builder.gray
}

// BgYellowBright returns a string where the background color is YellowBright.
func BgYellowBright(str ...string) string {
	return rootBuilder.BgYellowBright(str...)
}

// WithBgYellowBright returns a Builder that generates strings where the background color is YellowBright,
// and further styles can be applied via chaining.
func WithBgYellowBright() *Builder {
	return rootBuilder.WithBgYellowBright()
}

// BgYellowBright returns a string where the background color is YellowBright, in addition to other styles from this builder.
func (builder *Builder) BgYellowBright(str ...string) string {
	return builder.WithBgYellowBright().applyStyle(str...)
}

// WithBgYellowBright returns a Builder that generates strings where the background color is YellowBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithBgYellowBright() *Builder {
	if builder.bgYellowBright == nil {
		builder.bgYellowBright = createBuilder(builder, ansistyles.BgYellowBright.Open, ansistyles.BgYellowBright.Close)
	}
	return builder.bgYellowBright
}

// BgBlueBright returns a string where the background color is BlueBright.
func BgBlueBright(str ...string) string {
	return rootBuilder.BgBlueBright(str...)
}

// WithBgBlueBright returns a Builder that generates strings where the background color is BlueBright,
// and further styles can be applied via chaining.
func WithBgBlueBright() *Builder {
	return rootBuilder.WithBgBlueBright()
}

// BgBlueBright returns a string where the background color is BlueBright, in addition to other styles from this builder.
func (builder *Builder) BgBlueBright(str ...string) string {
	return builder.WithBgBlueBright().applyStyle(str...)
}

// WithBgBlueBright returns a Builder that generates strings where the background color is BlueBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithBgBlueBright() *Builder {
	if builder.bgBlueBright == nil {
		builder.bgBlueBright = createBuilder(builder, ansistyles.BgBlueBright.Open, ansistyles.BgBlueBright.Close)
	}
	return builder.bgBlueBright
}

// BgCyanBright returns a string where the background color is CyanBright.
func BgCyanBright(str ...string) string {
	return rootBuilder.BgCyanBright(str...)
}

// WithBgCyanBright returns a Builder that generates strings where the background color is CyanBright,
// and further styles can be applied via chaining.
func WithBgCyanBright() *Builder {
	return rootBuilder.WithBgCyanBright()
}

// BgCyanBright returns a string where the background color is CyanBright, in addition to other styles from this builder.
func (builder *Builder) BgCyanBright(str ...string) string {
	return builder.WithBgCyanBright().applyStyle(str...)
}

// WithBgCyanBright returns a Builder that generates strings where the background color is CyanBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithBgCyanBright() *Builder {
	if builder.bgCyanBright == nil {
		builder.bgCyanBright = createBuilder(builder, ansistyles.BgCyanBright.Open, ansistyles.BgCyanBright.Close)
	}
	return builder.bgCyanBright
}

// BgWhiteBright returns a string where the background color is WhiteBright.
func BgWhiteBright(str ...string) string {
	return rootBuilder.BgWhiteBright(str...)
}

// WithBgWhiteBright returns a Builder that generates strings where the background color is WhiteBright,
// and further styles can be applied via chaining.
func WithBgWhiteBright() *Builder {
	return rootBuilder.WithBgWhiteBright()
}

// BgWhiteBright returns a string where the background color is WhiteBright, in addition to other styles from this builder.
func (builder *Builder) BgWhiteBright(str ...string) string {
	return builder.WithBgWhiteBright().applyStyle(str...)
}

// WithBgWhiteBright returns a Builder that generates strings where the background color is WhiteBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithBgWhiteBright() *Builder {
	if builder.bgWhiteBright == nil {
		builder.bgWhiteBright = createBuilder(builder, ansistyles.BgWhiteBright.Open, ansistyles.BgWhiteBright.Close)
	}
	return builder.bgWhiteBright
}

// BgBlack returns a string where the background color is Black.
func BgBlack(str ...string) string {
	return rootBuilder.BgBlack(str...)
}

// WithBgBlack returns a Builder that generates strings where the background color is Black,
// and further styles can be applied via chaining.
func WithBgBlack() *Builder {
	return rootBuilder.WithBgBlack()
}

// BgBlack returns a string where the background color is Black, in addition to other styles from this builder.
func (builder *Builder) BgBlack(str ...string) string {
	return builder.WithBgBlack().applyStyle(str...)
}

// WithBgBlack returns a Builder that generates strings where the background color is Black,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithBgBlack() *Builder {
	if builder.bgBlack == nil {
		builder.bgBlack = createBuilder(builder, ansistyles.BgBlack.Open, ansistyles.BgBlack.Close)
	}
	return builder.bgBlack
}

// BgGreen returns a string where the background color is Green.
func BgGreen(str ...string) string {
	return rootBuilder.BgGreen(str...)
}

// WithBgGreen returns a Builder that generates strings where the background color is Green,
// and further styles can be applied via chaining.
func WithBgGreen() *Builder {
	return rootBuilder.WithBgGreen()
}

// BgGreen returns a string where the background color is Green, in addition to other styles from this builder.
func (builder *Builder) BgGreen(str ...string) string {
	return builder.WithBgGreen().applyStyle(str...)
}

// WithBgGreen returns a Builder that generates strings where the background color is Green,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithBgGreen() *Builder {
	if builder.bgGreen == nil {
		builder.bgGreen = createBuilder(builder, ansistyles.BgGreen.Open, ansistyles.BgGreen.Close)
	}
	return builder.bgGreen
}

// BgCyan returns a string where the background color is Cyan.
func BgCyan(str ...string) string {
	return rootBuilder.BgCyan(str...)
}

// WithBgCyan returns a Builder that generates strings where the background color is Cyan,
// and further styles can be applied via chaining.
func WithBgCyan() *Builder {
	return rootBuilder.WithBgCyan()
}

// BgCyan returns a string where the background color is Cyan, in addition to other styles from this builder.
func (builder *Builder) BgCyan(str ...string) string {
	return builder.WithBgCyan().applyStyle(str...)
}

// WithBgCyan returns a Builder that generates strings where the background color is Cyan,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithBgCyan() *Builder {
	if builder.bgCyan == nil {
		builder.bgCyan = createBuilder(builder, ansistyles.BgCyan.Open, ansistyles.BgCyan.Close)
	}
	return builder.bgCyan
}

// BgWhite returns a string where the background color is White.
func BgWhite(str ...string) string {
	return rootBuilder.BgWhite(str...)
}

// WithBgWhite returns a Builder that generates strings where the background color is White,
// and further styles can be applied via chaining.
func WithBgWhite() *Builder {
	return rootBuilder.WithBgWhite()
}

// BgWhite returns a string where the background color is White, in addition to other styles from this builder.
func (builder *Builder) BgWhite(str ...string) string {
	return builder.WithBgWhite().applyStyle(str...)
}

// WithBgWhite returns a Builder that generates strings where the background color is White,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithBgWhite() *Builder {
	if builder.bgWhite == nil {
		builder.bgWhite = createBuilder(builder, ansistyles.BgWhite.Open, ansistyles.BgWhite.Close)
	}
	return builder.bgWhite
}

// BgGrey returns a string where the background color is Grey.
func BgGrey(str ...string) string {
	return rootBuilder.BgGrey(str...)
}

// WithBgGrey returns a Builder that generates strings where the background color is Grey,
// and further styles can be applied via chaining.
func WithBgGrey() *Builder {
	return rootBuilder.WithBgGrey()
}

// BgGrey returns a string where the background color is Grey, in addition to other styles from this builder.
func (builder *Builder) BgGrey(str ...string) string {
	return builder.WithBgGrey().applyStyle(str...)
}

// WithBgGrey returns a Builder that generates strings where the background color is Grey,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithBgGrey() *Builder {
	if builder.bgGrey == nil {
		builder.bgGrey = createBuilder(builder, ansistyles.BgGrey.Open, ansistyles.BgGrey.Close)
	}
	return builder.bgGrey
}

// BgYellow returns a string where the background color is Yellow.
func BgYellow(str ...string) string {
	return rootBuilder.BgYellow(str...)
}

// WithBgYellow returns a Builder that generates strings where the background color is Yellow,
// and further styles can be applied via chaining.
func WithBgYellow() *Builder {
	return rootBuilder.WithBgYellow()
}

// BgYellow returns a string where the background color is Yellow, in addition to other styles from this builder.
func (builder *Builder) BgYellow(str ...string) string {
	return builder.WithBgYellow().applyStyle(str...)
}

// WithBgYellow returns a Builder that generates strings where the background color is Yellow,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithBgYellow() *Builder {
	if builder.bgYellow == nil {
		builder.bgYellow = createBuilder(builder, ansistyles.BgYellow.Open, ansistyles.BgYellow.Close)
	}
	return builder.bgYellow
}

// BgMagenta returns a string where the background color is Magenta.
func BgMagenta(str ...string) string {
	return rootBuilder.BgMagenta(str...)
}

// WithBgMagenta returns a Builder that generates strings where the background color is Magenta,
// and further styles can be applied via chaining.
func WithBgMagenta() *Builder {
	return rootBuilder.WithBgMagenta()
}

// BgMagenta returns a string where the background color is Magenta, in addition to other styles from this builder.
func (builder *Builder) BgMagenta(str ...string) string {
	return builder.WithBgMagenta().applyStyle(str...)
}

// WithBgMagenta returns a Builder that generates strings where the background color is Magenta,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithBgMagenta() *Builder {
	if builder.bgMagenta == nil {
		builder.bgMagenta = createBuilder(builder, ansistyles.BgMagenta.Open, ansistyles.BgMagenta.Close)
	}
	return builder.bgMagenta
}

// BgGray returns a string where the background color is Gray.
func BgGray(str ...string) string {
	return rootBuilder.BgGray(str...)
}

// WithBgGray returns a Builder that generates strings where the background color is Gray,
// and further styles can be applied via chaining.
func WithBgGray() *Builder {
	return rootBuilder.WithBgGray()
}

// BgGray returns a string where the background color is Gray, in addition to other styles from this builder.
func (builder *Builder) BgGray(str ...string) string {
	return builder.WithBgGray().applyStyle(str...)
}

// WithBgGray returns a Builder that generates strings where the background color is Gray,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithBgGray() *Builder {
	if builder.bgGray == nil {
		builder.bgGray = createBuilder(builder, ansistyles.BgGray.Open, ansistyles.BgGray.Close)
	}
	return builder.bgGray
}

// BgBlue returns a string where the background color is Blue.
func BgBlue(str ...string) string {
	return rootBuilder.BgBlue(str...)
}

// WithBgBlue returns a Builder that generates strings where the background color is Blue,
// and further styles can be applied via chaining.
func WithBgBlue() *Builder {
	return rootBuilder.WithBgBlue()
}

// BgBlue returns a string where the background color is Blue, in addition to other styles from this builder.
func (builder *Builder) BgBlue(str ...string) string {
	return builder.WithBgBlue().applyStyle(str...)
}

// WithBgBlue returns a Builder that generates strings where the background color is Blue,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithBgBlue() *Builder {
	if builder.bgBlue == nil {
		builder.bgBlue = createBuilder(builder, ansistyles.BgBlue.Open, ansistyles.BgBlue.Close)
	}
	return builder.bgBlue
}

// BgGreenBright returns a string where the background color is GreenBright.
func BgGreenBright(str ...string) string {
	return rootBuilder.BgGreenBright(str...)
}

// WithBgGreenBright returns a Builder that generates strings where the background color is GreenBright,
// and further styles can be applied via chaining.
func WithBgGreenBright() *Builder {
	return rootBuilder.WithBgGreenBright()
}

// BgGreenBright returns a string where the background color is GreenBright, in addition to other styles from this builder.
func (builder *Builder) BgGreenBright(str ...string) string {
	return builder.WithBgGreenBright().applyStyle(str...)
}

// WithBgGreenBright returns a Builder that generates strings where the background color is GreenBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithBgGreenBright() *Builder {
	if builder.bgGreenBright == nil {
		builder.bgGreenBright = createBuilder(builder, ansistyles.BgGreenBright.Open, ansistyles.BgGreenBright.Close)
	}
	return builder.bgGreenBright
}

// BgMagentaBright returns a string where the background color is MagentaBright.
func BgMagentaBright(str ...string) string {
	return rootBuilder.BgMagentaBright(str...)
}

// WithBgMagentaBright returns a Builder that generates strings where the background color is MagentaBright,
// and further styles can be applied via chaining.
func WithBgMagentaBright() *Builder {
	return rootBuilder.WithBgMagentaBright()
}

// BgMagentaBright returns a string where the background color is MagentaBright, in addition to other styles from this builder.
func (builder *Builder) BgMagentaBright(str ...string) string {
	return builder.WithBgMagentaBright().applyStyle(str...)
}

// WithBgMagentaBright returns a Builder that generates strings where the background color is MagentaBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithBgMagentaBright() *Builder {
	if builder.bgMagentaBright == nil {
		builder.bgMagentaBright = createBuilder(builder, ansistyles.BgMagentaBright.Open, ansistyles.BgMagentaBright.Close)
	}
	return builder.bgMagentaBright
}

// BgRed returns a string where the background color is Red.
func BgRed(str ...string) string {
	return rootBuilder.BgRed(str...)
}

// WithBgRed returns a Builder that generates strings where the background color is Red,
// and further styles can be applied via chaining.
func WithBgRed() *Builder {
	return rootBuilder.WithBgRed()
}

// BgRed returns a string where the background color is Red, in addition to other styles from this builder.
func (builder *Builder) BgRed(str ...string) string {
	return builder.WithBgRed().applyStyle(str...)
}

// WithBgRed returns a Builder that generates strings where the background color is Red,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithBgRed() *Builder {
	if builder.bgRed == nil {
		builder.bgRed = createBuilder(builder, ansistyles.BgRed.Open, ansistyles.BgRed.Close)
	}
	return builder.bgRed
}

// BgBlackBright returns a string where the background color is BlackBright.
func BgBlackBright(str ...string) string {
	return rootBuilder.BgBlackBright(str...)
}

// WithBgBlackBright returns a Builder that generates strings where the background color is BlackBright,
// and further styles can be applied via chaining.
func WithBgBlackBright() *Builder {
	return rootBuilder.WithBgBlackBright()
}

// BgBlackBright returns a string where the background color is BlackBright, in addition to other styles from this builder.
func (builder *Builder) BgBlackBright(str ...string) string {
	return builder.WithBgBlackBright().applyStyle(str...)
}

// WithBgBlackBright returns a Builder that generates strings where the background color is BlackBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithBgBlackBright() *Builder {
	if builder.bgBlackBright == nil {
		builder.bgBlackBright = createBuilder(builder, ansistyles.BgBlackBright.Open, ansistyles.BgBlackBright.Close)
	}
	return builder.bgBlackBright
}

// BgRedBright returns a string where the background color is RedBright.
func BgRedBright(str ...string) string {
	return rootBuilder.BgRedBright(str...)
}

// WithBgRedBright returns a Builder that generates strings where the background color is RedBright,
// and further styles can be applied via chaining.
func WithBgRedBright() *Builder {
	return rootBuilder.WithBgRedBright()
}

// BgRedBright returns a string where the background color is RedBright, in addition to other styles from this builder.
func (builder *Builder) BgRedBright(str ...string) string {
	return builder.WithBgRedBright().applyStyle(str...)
}

// WithBgRedBright returns a Builder that generates strings where the background color is RedBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithBgRedBright() *Builder {
	if builder.bgRedBright == nil {
		builder.bgRedBright = createBuilder(builder, ansistyles.BgRedBright.Open, ansistyles.BgRedBright.Close)
	}
	return builder.bgRedBright
}

// Reset returns a string with the reset modifier.
func Reset(str ...string) string {
	return rootBuilder.Reset(str...)
}

// WithReset returns a Builder that generates strings with the reset modifier,
// and further styles can be applied via chaining.
func WithReset() *Builder {
	return rootBuilder.WithReset()
}

// Reset returns a string with the reset modifier, in addition to other styles from this builder.
func (builder *Builder) Reset(str ...string) string {
	return builder.WithReset().applyStyle(str...)
}

// WithReset returns a Builder that generates strings with the reset modifier,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithReset() *Builder {
	if builder.reset == nil {
		builder.reset = createBuilder(builder, ansistyles.Reset.Open, ansistyles.Reset.Close)
	}
	return builder.reset
}

// Bold returns a string with the bold modifier.
func Bold(str ...string) string {
	return rootBuilder.Bold(str...)
}

// WithBold returns a Builder that generates strings with the bold modifier,
// and further styles can be applied via chaining.
func WithBold() *Builder {
	return rootBuilder.WithBold()
}

// Bold returns a string with the bold modifier, in addition to other styles from this builder.
func (builder *Builder) Bold(str ...string) string {
	return builder.WithBold().applyStyle(str...)
}

// WithBold returns a Builder that generates strings with the bold modifier,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithBold() *Builder {
	if builder.bold == nil {
		builder.bold = createBuilder(builder, ansistyles.Bold.Open, ansistyles.Bold.Close)
	}
	return builder.bold
}

// Italic returns a string with the italic modifier.
func Italic(str ...string) string {
	return rootBuilder.Italic(str...)
}

// WithItalic returns a Builder that generates strings with the italic modifier,
// and further styles can be applied via chaining.
func WithItalic() *Builder {
	return rootBuilder.WithItalic()
}

// Italic returns a string with the italic modifier, in addition to other styles from this builder.
func (builder *Builder) Italic(str ...string) string {
	return builder.WithItalic().applyStyle(str...)
}

// WithItalic returns a Builder that generates strings with the italic modifier,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithItalic() *Builder {
	if builder.italic == nil {
		builder.italic = createBuilder(builder, ansistyles.Italic.Open, ansistyles.Italic.Close)
	}
	return builder.italic
}

// Underline returns a string with the underline modifier.
func Underline(str ...string) string {
	return rootBuilder.Underline(str...)
}

// WithUnderline returns a Builder that generates strings with the underline modifier,
// and further styles can be applied via chaining.
func WithUnderline() *Builder {
	return rootBuilder.WithUnderline()
}

// Underline returns a string with the underline modifier, in addition to other styles from this builder.
func (builder *Builder) Underline(str ...string) string {
	return builder.WithUnderline().applyStyle(str...)
}

// WithUnderline returns a Builder that generates strings with the underline modifier,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithUnderline() *Builder {
	if builder.underline == nil {
		builder.underline = createBuilder(builder, ansistyles.Underline.Open, ansistyles.Underline.Close)
	}
	return builder.underline
}

// Hidden returns a string with the hidden modifier.
func Hidden(str ...string) string {
	return rootBuilder.Hidden(str...)
}

// WithHidden returns a Builder that generates strings with the hidden modifier,
// and further styles can be applied via chaining.
func WithHidden() *Builder {
	return rootBuilder.WithHidden()
}

// Hidden returns a string with the hidden modifier, in addition to other styles from this builder.
func (builder *Builder) Hidden(str ...string) string {
	return builder.WithHidden().applyStyle(str...)
}

// WithHidden returns a Builder that generates strings with the hidden modifier,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithHidden() *Builder {
	if builder.hidden == nil {
		builder.hidden = createBuilder(builder, ansistyles.Hidden.Open, ansistyles.Hidden.Close)
	}
	return builder.hidden
}

// Dim returns a string with the dim modifier.
func Dim(str ...string) string {
	return rootBuilder.Dim(str...)
}

// WithDim returns a Builder that generates strings with the dim modifier,
// and further styles can be applied via chaining.
func WithDim() *Builder {
	return rootBuilder.WithDim()
}

// Dim returns a string with the dim modifier, in addition to other styles from this builder.
func (builder *Builder) Dim(str ...string) string {
	return builder.WithDim().applyStyle(str...)
}

// WithDim returns a Builder that generates strings with the dim modifier,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithDim() *Builder {
	if builder.dim == nil {
		builder.dim = createBuilder(builder, ansistyles.Dim.Open, ansistyles.Dim.Close)
	}
	return builder.dim
}

// Overline returns a string with the overline modifier.
func Overline(str ...string) string {
	return rootBuilder.Overline(str...)
}

// WithOverline returns a Builder that generates strings with the overline modifier,
// and further styles can be applied via chaining.
func WithOverline() *Builder {
	return rootBuilder.WithOverline()
}

// Overline returns a string with the overline modifier, in addition to other styles from this builder.
func (builder *Builder) Overline(str ...string) string {
	return builder.WithOverline().applyStyle(str...)
}

// WithOverline returns a Builder that generates strings with the overline modifier,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithOverline() *Builder {
	if builder.overline == nil {
		builder.overline = createBuilder(builder, ansistyles.Overline.Open, ansistyles.Overline.Close)
	}
	return builder.overline
}

// Inverse returns a string with the inverse modifier.
func Inverse(str ...string) string {
	return rootBuilder.Inverse(str...)
}

// WithInverse returns a Builder that generates strings with the inverse modifier,
// and further styles can be applied via chaining.
func WithInverse() *Builder {
	return rootBuilder.WithInverse()
}

// Inverse returns a string with the inverse modifier, in addition to other styles from this builder.
func (builder *Builder) Inverse(str ...string) string {
	return builder.WithInverse().applyStyle(str...)
}

// WithInverse returns a Builder that generates strings with the inverse modifier,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithInverse() *Builder {
	if builder.inverse == nil {
		builder.inverse = createBuilder(builder, ansistyles.Inverse.Open, ansistyles.Inverse.Close)
	}
	return builder.inverse
}

// Strikethrough returns a string with the strikethrough modifier.
func Strikethrough(str ...string) string {
	return rootBuilder.Strikethrough(str...)
}

// WithStrikethrough returns a Builder that generates strings with the strikethrough modifier,
// and further styles can be applied via chaining.
func WithStrikethrough() *Builder {
	return rootBuilder.WithStrikethrough()
}

// Strikethrough returns a string with the strikethrough modifier, in addition to other styles from this builder.
func (builder *Builder) Strikethrough(str ...string) string {
	return builder.WithStrikethrough().applyStyle(str...)
}

// WithStrikethrough returns a Builder that generates strings with the strikethrough modifier,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WithStrikethrough() *Builder {
	if builder.strikethrough == nil {
		builder.strikethrough = createBuilder(builder, ansistyles.Strikethrough.Open, ansistyles.Strikethrough.Close)
	}
	return builder.strikethrough
}

type generatedBuilders struct {
	bgWhiteBright   *Builder
	underline       *Builder
	strikethrough   *Builder
	green           *Builder
	blackBright     *Builder
	redBright       *Builder
	bgWhite         *Builder
	cyan            *Builder
	bgYellowBright  *Builder
	bgGray          *Builder
	bgCyanBright    *Builder
	red             *Builder
	white           *Builder
	whiteBright     *Builder
	bgBlack         *Builder
	bgRed           *Builder
	bgGreen         *Builder
	italic          *Builder
	hidden          *Builder
	magenta         *Builder
	greenBright     *Builder
	yellowBright    *Builder
	bgGreenBright   *Builder
	bgBlueBright    *Builder
	bold            *Builder
	inverse         *Builder
	black           *Builder
	magentaBright   *Builder
	gray            *Builder
	bgBlue          *Builder
	yellow          *Builder
	blue            *Builder
	bgYellow        *Builder
	bgMagentaBright *Builder
	bgGrey          *Builder
	dim             *Builder
	cyanBright      *Builder
	grey            *Builder
	bgMagenta       *Builder
	reset           *Builder
	overline        *Builder
	blueBright      *Builder
	bgCyan          *Builder
	bgBlackBright   *Builder
	bgRedBright     *Builder
}
