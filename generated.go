package gawk

// This file was generated.  Don't edit it directly.

import (
	"github.com/jwalton/go-ansistyles"
)

// Black returns a string where the color is black.
func Black(str ...string) string {
	return rootBuilder.Black(str...)
}

// BlackAnd returns a Builder that generates strings where the color is black,
// and further styles can be applied via chaining.
func BlackAnd() *Builder {
	return rootBuilder.BlackAnd()
}

// Black returns a string where the color is black, in addition to other styles from this builder.
func (builder *Builder) Black(str ...string) string {
	return builder.BlackAnd().applyStyle(str...)
}

// BlackAnd returns a Builder that generates strings where the color is black,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) BlackAnd() *Builder {
	if builder.black == nil {
		builder.black = createBuilder(builder, ansistyles.Black.Open, ansistyles.Black.Close)
	}
	return builder.black
}

// Green returns a string where the color is green.
func Green(str ...string) string {
	return rootBuilder.Green(str...)
}

// GreenAnd returns a Builder that generates strings where the color is green,
// and further styles can be applied via chaining.
func GreenAnd() *Builder {
	return rootBuilder.GreenAnd()
}

// Green returns a string where the color is green, in addition to other styles from this builder.
func (builder *Builder) Green(str ...string) string {
	return builder.GreenAnd().applyStyle(str...)
}

// GreenAnd returns a Builder that generates strings where the color is green,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) GreenAnd() *Builder {
	if builder.green == nil {
		builder.green = createBuilder(builder, ansistyles.Green.Open, ansistyles.Green.Close)
	}
	return builder.green
}

// Magenta returns a string where the color is magenta.
func Magenta(str ...string) string {
	return rootBuilder.Magenta(str...)
}

// MagentaAnd returns a Builder that generates strings where the color is magenta,
// and further styles can be applied via chaining.
func MagentaAnd() *Builder {
	return rootBuilder.MagentaAnd()
}

// Magenta returns a string where the color is magenta, in addition to other styles from this builder.
func (builder *Builder) Magenta(str ...string) string {
	return builder.MagentaAnd().applyStyle(str...)
}

// MagentaAnd returns a Builder that generates strings where the color is magenta,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) MagentaAnd() *Builder {
	if builder.magenta == nil {
		builder.magenta = createBuilder(builder, ansistyles.Magenta.Open, ansistyles.Magenta.Close)
	}
	return builder.magenta
}

// Cyan returns a string where the color is cyan.
func Cyan(str ...string) string {
	return rootBuilder.Cyan(str...)
}

// CyanAnd returns a Builder that generates strings where the color is cyan,
// and further styles can be applied via chaining.
func CyanAnd() *Builder {
	return rootBuilder.CyanAnd()
}

// Cyan returns a string where the color is cyan, in addition to other styles from this builder.
func (builder *Builder) Cyan(str ...string) string {
	return builder.CyanAnd().applyStyle(str...)
}

// CyanAnd returns a Builder that generates strings where the color is cyan,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) CyanAnd() *Builder {
	if builder.cyan == nil {
		builder.cyan = createBuilder(builder, ansistyles.Cyan.Open, ansistyles.Cyan.Close)
	}
	return builder.cyan
}

// Yellow returns a string where the color is yellow.
func Yellow(str ...string) string {
	return rootBuilder.Yellow(str...)
}

// YellowAnd returns a Builder that generates strings where the color is yellow,
// and further styles can be applied via chaining.
func YellowAnd() *Builder {
	return rootBuilder.YellowAnd()
}

// Yellow returns a string where the color is yellow, in addition to other styles from this builder.
func (builder *Builder) Yellow(str ...string) string {
	return builder.YellowAnd().applyStyle(str...)
}

// YellowAnd returns a Builder that generates strings where the color is yellow,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) YellowAnd() *Builder {
	if builder.yellow == nil {
		builder.yellow = createBuilder(builder, ansistyles.Yellow.Open, ansistyles.Yellow.Close)
	}
	return builder.yellow
}

// Blue returns a string where the color is blue.
func Blue(str ...string) string {
	return rootBuilder.Blue(str...)
}

// BlueAnd returns a Builder that generates strings where the color is blue,
// and further styles can be applied via chaining.
func BlueAnd() *Builder {
	return rootBuilder.BlueAnd()
}

// Blue returns a string where the color is blue, in addition to other styles from this builder.
func (builder *Builder) Blue(str ...string) string {
	return builder.BlueAnd().applyStyle(str...)
}

// BlueAnd returns a Builder that generates strings where the color is blue,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) BlueAnd() *Builder {
	if builder.blue == nil {
		builder.blue = createBuilder(builder, ansistyles.Blue.Open, ansistyles.Blue.Close)
	}
	return builder.blue
}

// BlackBright returns a string where the color is blackBright.
func BlackBright(str ...string) string {
	return rootBuilder.BlackBright(str...)
}

// BlackBrightAnd returns a Builder that generates strings where the color is blackBright,
// and further styles can be applied via chaining.
func BlackBrightAnd() *Builder {
	return rootBuilder.BlackBrightAnd()
}

// BlackBright returns a string where the color is blackBright, in addition to other styles from this builder.
func (builder *Builder) BlackBright(str ...string) string {
	return builder.BlackBrightAnd().applyStyle(str...)
}

// BlackBrightAnd returns a Builder that generates strings where the color is blackBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) BlackBrightAnd() *Builder {
	if builder.blackBright == nil {
		builder.blackBright = createBuilder(builder, ansistyles.BlackBright.Open, ansistyles.BlackBright.Close)
	}
	return builder.blackBright
}

// GreenBright returns a string where the color is greenBright.
func GreenBright(str ...string) string {
	return rootBuilder.GreenBright(str...)
}

// GreenBrightAnd returns a Builder that generates strings where the color is greenBright,
// and further styles can be applied via chaining.
func GreenBrightAnd() *Builder {
	return rootBuilder.GreenBrightAnd()
}

// GreenBright returns a string where the color is greenBright, in addition to other styles from this builder.
func (builder *Builder) GreenBright(str ...string) string {
	return builder.GreenBrightAnd().applyStyle(str...)
}

// GreenBrightAnd returns a Builder that generates strings where the color is greenBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) GreenBrightAnd() *Builder {
	if builder.greenBright == nil {
		builder.greenBright = createBuilder(builder, ansistyles.GreenBright.Open, ansistyles.GreenBright.Close)
	}
	return builder.greenBright
}

// WhiteBright returns a string where the color is whiteBright.
func WhiteBright(str ...string) string {
	return rootBuilder.WhiteBright(str...)
}

// WhiteBrightAnd returns a Builder that generates strings where the color is whiteBright,
// and further styles can be applied via chaining.
func WhiteBrightAnd() *Builder {
	return rootBuilder.WhiteBrightAnd()
}

// WhiteBright returns a string where the color is whiteBright, in addition to other styles from this builder.
func (builder *Builder) WhiteBright(str ...string) string {
	return builder.WhiteBrightAnd().applyStyle(str...)
}

// WhiteBrightAnd returns a Builder that generates strings where the color is whiteBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WhiteBrightAnd() *Builder {
	if builder.whiteBright == nil {
		builder.whiteBright = createBuilder(builder, ansistyles.WhiteBright.Open, ansistyles.WhiteBright.Close)
	}
	return builder.whiteBright
}

// Grey returns a string where the color is grey.
func Grey(str ...string) string {
	return rootBuilder.Grey(str...)
}

// GreyAnd returns a Builder that generates strings where the color is grey,
// and further styles can be applied via chaining.
func GreyAnd() *Builder {
	return rootBuilder.GreyAnd()
}

// Grey returns a string where the color is grey, in addition to other styles from this builder.
func (builder *Builder) Grey(str ...string) string {
	return builder.GreyAnd().applyStyle(str...)
}

// GreyAnd returns a Builder that generates strings where the color is grey,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) GreyAnd() *Builder {
	if builder.grey == nil {
		builder.grey = createBuilder(builder, ansistyles.Grey.Open, ansistyles.Grey.Close)
	}
	return builder.grey
}

// Gray returns a string where the color is gray.
func Gray(str ...string) string {
	return rootBuilder.Gray(str...)
}

// GrayAnd returns a Builder that generates strings where the color is gray,
// and further styles can be applied via chaining.
func GrayAnd() *Builder {
	return rootBuilder.GrayAnd()
}

// Gray returns a string where the color is gray, in addition to other styles from this builder.
func (builder *Builder) Gray(str ...string) string {
	return builder.GrayAnd().applyStyle(str...)
}

// GrayAnd returns a Builder that generates strings where the color is gray,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) GrayAnd() *Builder {
	if builder.gray == nil {
		builder.gray = createBuilder(builder, ansistyles.Gray.Open, ansistyles.Gray.Close)
	}
	return builder.gray
}

// Red returns a string where the color is red.
func Red(str ...string) string {
	return rootBuilder.Red(str...)
}

// RedAnd returns a Builder that generates strings where the color is red,
// and further styles can be applied via chaining.
func RedAnd() *Builder {
	return rootBuilder.RedAnd()
}

// Red returns a string where the color is red, in addition to other styles from this builder.
func (builder *Builder) Red(str ...string) string {
	return builder.RedAnd().applyStyle(str...)
}

// RedAnd returns a Builder that generates strings where the color is red,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) RedAnd() *Builder {
	if builder.red == nil {
		builder.red = createBuilder(builder, ansistyles.Red.Open, ansistyles.Red.Close)
	}
	return builder.red
}

// RedBright returns a string where the color is redBright.
func RedBright(str ...string) string {
	return rootBuilder.RedBright(str...)
}

// RedBrightAnd returns a Builder that generates strings where the color is redBright,
// and further styles can be applied via chaining.
func RedBrightAnd() *Builder {
	return rootBuilder.RedBrightAnd()
}

// RedBright returns a string where the color is redBright, in addition to other styles from this builder.
func (builder *Builder) RedBright(str ...string) string {
	return builder.RedBrightAnd().applyStyle(str...)
}

// RedBrightAnd returns a Builder that generates strings where the color is redBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) RedBrightAnd() *Builder {
	if builder.redBright == nil {
		builder.redBright = createBuilder(builder, ansistyles.RedBright.Open, ansistyles.RedBright.Close)
	}
	return builder.redBright
}

// BlueBright returns a string where the color is blueBright.
func BlueBright(str ...string) string {
	return rootBuilder.BlueBright(str...)
}

// BlueBrightAnd returns a Builder that generates strings where the color is blueBright,
// and further styles can be applied via chaining.
func BlueBrightAnd() *Builder {
	return rootBuilder.BlueBrightAnd()
}

// BlueBright returns a string where the color is blueBright, in addition to other styles from this builder.
func (builder *Builder) BlueBright(str ...string) string {
	return builder.BlueBrightAnd().applyStyle(str...)
}

// BlueBrightAnd returns a Builder that generates strings where the color is blueBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) BlueBrightAnd() *Builder {
	if builder.blueBright == nil {
		builder.blueBright = createBuilder(builder, ansistyles.BlueBright.Open, ansistyles.BlueBright.Close)
	}
	return builder.blueBright
}

// MagentaBright returns a string where the color is magentaBright.
func MagentaBright(str ...string) string {
	return rootBuilder.MagentaBright(str...)
}

// MagentaBrightAnd returns a Builder that generates strings where the color is magentaBright,
// and further styles can be applied via chaining.
func MagentaBrightAnd() *Builder {
	return rootBuilder.MagentaBrightAnd()
}

// MagentaBright returns a string where the color is magentaBright, in addition to other styles from this builder.
func (builder *Builder) MagentaBright(str ...string) string {
	return builder.MagentaBrightAnd().applyStyle(str...)
}

// MagentaBrightAnd returns a Builder that generates strings where the color is magentaBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) MagentaBrightAnd() *Builder {
	if builder.magentaBright == nil {
		builder.magentaBright = createBuilder(builder, ansistyles.MagentaBright.Open, ansistyles.MagentaBright.Close)
	}
	return builder.magentaBright
}

// CyanBright returns a string where the color is cyanBright.
func CyanBright(str ...string) string {
	return rootBuilder.CyanBright(str...)
}

// CyanBrightAnd returns a Builder that generates strings where the color is cyanBright,
// and further styles can be applied via chaining.
func CyanBrightAnd() *Builder {
	return rootBuilder.CyanBrightAnd()
}

// CyanBright returns a string where the color is cyanBright, in addition to other styles from this builder.
func (builder *Builder) CyanBright(str ...string) string {
	return builder.CyanBrightAnd().applyStyle(str...)
}

// CyanBrightAnd returns a Builder that generates strings where the color is cyanBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) CyanBrightAnd() *Builder {
	if builder.cyanBright == nil {
		builder.cyanBright = createBuilder(builder, ansistyles.CyanBright.Open, ansistyles.CyanBright.Close)
	}
	return builder.cyanBright
}

// White returns a string where the color is white.
func White(str ...string) string {
	return rootBuilder.White(str...)
}

// WhiteAnd returns a Builder that generates strings where the color is white,
// and further styles can be applied via chaining.
func WhiteAnd() *Builder {
	return rootBuilder.WhiteAnd()
}

// White returns a string where the color is white, in addition to other styles from this builder.
func (builder *Builder) White(str ...string) string {
	return builder.WhiteAnd().applyStyle(str...)
}

// WhiteAnd returns a Builder that generates strings where the color is white,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) WhiteAnd() *Builder {
	if builder.white == nil {
		builder.white = createBuilder(builder, ansistyles.White.Open, ansistyles.White.Close)
	}
	return builder.white
}

// YellowBright returns a string where the color is yellowBright.
func YellowBright(str ...string) string {
	return rootBuilder.YellowBright(str...)
}

// YellowBrightAnd returns a Builder that generates strings where the color is yellowBright,
// and further styles can be applied via chaining.
func YellowBrightAnd() *Builder {
	return rootBuilder.YellowBrightAnd()
}

// YellowBright returns a string where the color is yellowBright, in addition to other styles from this builder.
func (builder *Builder) YellowBright(str ...string) string {
	return builder.YellowBrightAnd().applyStyle(str...)
}

// YellowBrightAnd returns a Builder that generates strings where the color is yellowBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) YellowBrightAnd() *Builder {
	if builder.yellowBright == nil {
		builder.yellowBright = createBuilder(builder, ansistyles.YellowBright.Open, ansistyles.YellowBright.Close)
	}
	return builder.yellowBright
}

// BgCyan returns a string where the background color is Cyan.
func BgCyan(str ...string) string {
	return rootBuilder.BgCyan(str...)
}

// BgCyanAnd returns a Builder that generates strings where the background color is Cyan,
// and further styles can be applied via chaining.
func BgCyanAnd() *Builder {
	return rootBuilder.BgCyanAnd()
}

// BgCyan returns a string where the background color is Cyan, in addition to other styles from this builder.
func (builder *Builder) BgCyan(str ...string) string {
	return builder.BgCyanAnd().applyStyle(str...)
}

// BgCyanAnd returns a Builder that generates strings where the background color is Cyan,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) BgCyanAnd() *Builder {
	if builder.bgCyan == nil {
		builder.bgCyan = createBuilder(builder, ansistyles.BgCyan.Open, ansistyles.BgCyan.Close)
	}
	return builder.bgCyan
}

// BgBlueBright returns a string where the background color is BlueBright.
func BgBlueBright(str ...string) string {
	return rootBuilder.BgBlueBright(str...)
}

// BgBlueBrightAnd returns a Builder that generates strings where the background color is BlueBright,
// and further styles can be applied via chaining.
func BgBlueBrightAnd() *Builder {
	return rootBuilder.BgBlueBrightAnd()
}

// BgBlueBright returns a string where the background color is BlueBright, in addition to other styles from this builder.
func (builder *Builder) BgBlueBright(str ...string) string {
	return builder.BgBlueBrightAnd().applyStyle(str...)
}

// BgBlueBrightAnd returns a Builder that generates strings where the background color is BlueBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) BgBlueBrightAnd() *Builder {
	if builder.bgBlueBright == nil {
		builder.bgBlueBright = createBuilder(builder, ansistyles.BgBlueBright.Open, ansistyles.BgBlueBright.Close)
	}
	return builder.bgBlueBright
}

// BgGrey returns a string where the background color is Grey.
func BgGrey(str ...string) string {
	return rootBuilder.BgGrey(str...)
}

// BgGreyAnd returns a Builder that generates strings where the background color is Grey,
// and further styles can be applied via chaining.
func BgGreyAnd() *Builder {
	return rootBuilder.BgGreyAnd()
}

// BgGrey returns a string where the background color is Grey, in addition to other styles from this builder.
func (builder *Builder) BgGrey(str ...string) string {
	return builder.BgGreyAnd().applyStyle(str...)
}

// BgGreyAnd returns a Builder that generates strings where the background color is Grey,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) BgGreyAnd() *Builder {
	if builder.bgGrey == nil {
		builder.bgGrey = createBuilder(builder, ansistyles.BgGrey.Open, ansistyles.BgGrey.Close)
	}
	return builder.bgGrey
}

// BgGreen returns a string where the background color is Green.
func BgGreen(str ...string) string {
	return rootBuilder.BgGreen(str...)
}

// BgGreenAnd returns a Builder that generates strings where the background color is Green,
// and further styles can be applied via chaining.
func BgGreenAnd() *Builder {
	return rootBuilder.BgGreenAnd()
}

// BgGreen returns a string where the background color is Green, in addition to other styles from this builder.
func (builder *Builder) BgGreen(str ...string) string {
	return builder.BgGreenAnd().applyStyle(str...)
}

// BgGreenAnd returns a Builder that generates strings where the background color is Green,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) BgGreenAnd() *Builder {
	if builder.bgGreen == nil {
		builder.bgGreen = createBuilder(builder, ansistyles.BgGreen.Open, ansistyles.BgGreen.Close)
	}
	return builder.bgGreen
}

// BgGreenBright returns a string where the background color is GreenBright.
func BgGreenBright(str ...string) string {
	return rootBuilder.BgGreenBright(str...)
}

// BgGreenBrightAnd returns a Builder that generates strings where the background color is GreenBright,
// and further styles can be applied via chaining.
func BgGreenBrightAnd() *Builder {
	return rootBuilder.BgGreenBrightAnd()
}

// BgGreenBright returns a string where the background color is GreenBright, in addition to other styles from this builder.
func (builder *Builder) BgGreenBright(str ...string) string {
	return builder.BgGreenBrightAnd().applyStyle(str...)
}

// BgGreenBrightAnd returns a Builder that generates strings where the background color is GreenBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) BgGreenBrightAnd() *Builder {
	if builder.bgGreenBright == nil {
		builder.bgGreenBright = createBuilder(builder, ansistyles.BgGreenBright.Open, ansistyles.BgGreenBright.Close)
	}
	return builder.bgGreenBright
}

// BgMagenta returns a string where the background color is Magenta.
func BgMagenta(str ...string) string {
	return rootBuilder.BgMagenta(str...)
}

// BgMagentaAnd returns a Builder that generates strings where the background color is Magenta,
// and further styles can be applied via chaining.
func BgMagentaAnd() *Builder {
	return rootBuilder.BgMagentaAnd()
}

// BgMagenta returns a string where the background color is Magenta, in addition to other styles from this builder.
func (builder *Builder) BgMagenta(str ...string) string {
	return builder.BgMagentaAnd().applyStyle(str...)
}

// BgMagentaAnd returns a Builder that generates strings where the background color is Magenta,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) BgMagentaAnd() *Builder {
	if builder.bgMagenta == nil {
		builder.bgMagenta = createBuilder(builder, ansistyles.BgMagenta.Open, ansistyles.BgMagenta.Close)
	}
	return builder.bgMagenta
}

// BgBlue returns a string where the background color is Blue.
func BgBlue(str ...string) string {
	return rootBuilder.BgBlue(str...)
}

// BgBlueAnd returns a Builder that generates strings where the background color is Blue,
// and further styles can be applied via chaining.
func BgBlueAnd() *Builder {
	return rootBuilder.BgBlueAnd()
}

// BgBlue returns a string where the background color is Blue, in addition to other styles from this builder.
func (builder *Builder) BgBlue(str ...string) string {
	return builder.BgBlueAnd().applyStyle(str...)
}

// BgBlueAnd returns a Builder that generates strings where the background color is Blue,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) BgBlueAnd() *Builder {
	if builder.bgBlue == nil {
		builder.bgBlue = createBuilder(builder, ansistyles.BgBlue.Open, ansistyles.BgBlue.Close)
	}
	return builder.bgBlue
}

// BgWhite returns a string where the background color is White.
func BgWhite(str ...string) string {
	return rootBuilder.BgWhite(str...)
}

// BgWhiteAnd returns a Builder that generates strings where the background color is White,
// and further styles can be applied via chaining.
func BgWhiteAnd() *Builder {
	return rootBuilder.BgWhiteAnd()
}

// BgWhite returns a string where the background color is White, in addition to other styles from this builder.
func (builder *Builder) BgWhite(str ...string) string {
	return builder.BgWhiteAnd().applyStyle(str...)
}

// BgWhiteAnd returns a Builder that generates strings where the background color is White,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) BgWhiteAnd() *Builder {
	if builder.bgWhite == nil {
		builder.bgWhite = createBuilder(builder, ansistyles.BgWhite.Open, ansistyles.BgWhite.Close)
	}
	return builder.bgWhite
}

// BgYellowBright returns a string where the background color is YellowBright.
func BgYellowBright(str ...string) string {
	return rootBuilder.BgYellowBright(str...)
}

// BgYellowBrightAnd returns a Builder that generates strings where the background color is YellowBright,
// and further styles can be applied via chaining.
func BgYellowBrightAnd() *Builder {
	return rootBuilder.BgYellowBrightAnd()
}

// BgYellowBright returns a string where the background color is YellowBright, in addition to other styles from this builder.
func (builder *Builder) BgYellowBright(str ...string) string {
	return builder.BgYellowBrightAnd().applyStyle(str...)
}

// BgYellowBrightAnd returns a Builder that generates strings where the background color is YellowBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) BgYellowBrightAnd() *Builder {
	if builder.bgYellowBright == nil {
		builder.bgYellowBright = createBuilder(builder, ansistyles.BgYellowBright.Open, ansistyles.BgYellowBright.Close)
	}
	return builder.bgYellowBright
}

// BgWhiteBright returns a string where the background color is WhiteBright.
func BgWhiteBright(str ...string) string {
	return rootBuilder.BgWhiteBright(str...)
}

// BgWhiteBrightAnd returns a Builder that generates strings where the background color is WhiteBright,
// and further styles can be applied via chaining.
func BgWhiteBrightAnd() *Builder {
	return rootBuilder.BgWhiteBrightAnd()
}

// BgWhiteBright returns a string where the background color is WhiteBright, in addition to other styles from this builder.
func (builder *Builder) BgWhiteBright(str ...string) string {
	return builder.BgWhiteBrightAnd().applyStyle(str...)
}

// BgWhiteBrightAnd returns a Builder that generates strings where the background color is WhiteBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) BgWhiteBrightAnd() *Builder {
	if builder.bgWhiteBright == nil {
		builder.bgWhiteBright = createBuilder(builder, ansistyles.BgWhiteBright.Open, ansistyles.BgWhiteBright.Close)
	}
	return builder.bgWhiteBright
}

// BgGray returns a string where the background color is Gray.
func BgGray(str ...string) string {
	return rootBuilder.BgGray(str...)
}

// BgGrayAnd returns a Builder that generates strings where the background color is Gray,
// and further styles can be applied via chaining.
func BgGrayAnd() *Builder {
	return rootBuilder.BgGrayAnd()
}

// BgGray returns a string where the background color is Gray, in addition to other styles from this builder.
func (builder *Builder) BgGray(str ...string) string {
	return builder.BgGrayAnd().applyStyle(str...)
}

// BgGrayAnd returns a Builder that generates strings where the background color is Gray,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) BgGrayAnd() *Builder {
	if builder.bgGray == nil {
		builder.bgGray = createBuilder(builder, ansistyles.BgGray.Open, ansistyles.BgGray.Close)
	}
	return builder.bgGray
}

// BgRed returns a string where the background color is Red.
func BgRed(str ...string) string {
	return rootBuilder.BgRed(str...)
}

// BgRedAnd returns a Builder that generates strings where the background color is Red,
// and further styles can be applied via chaining.
func BgRedAnd() *Builder {
	return rootBuilder.BgRedAnd()
}

// BgRed returns a string where the background color is Red, in addition to other styles from this builder.
func (builder *Builder) BgRed(str ...string) string {
	return builder.BgRedAnd().applyStyle(str...)
}

// BgRedAnd returns a Builder that generates strings where the background color is Red,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) BgRedAnd() *Builder {
	if builder.bgRed == nil {
		builder.bgRed = createBuilder(builder, ansistyles.BgRed.Open, ansistyles.BgRed.Close)
	}
	return builder.bgRed
}

// BgYellow returns a string where the background color is Yellow.
func BgYellow(str ...string) string {
	return rootBuilder.BgYellow(str...)
}

// BgYellowAnd returns a Builder that generates strings where the background color is Yellow,
// and further styles can be applied via chaining.
func BgYellowAnd() *Builder {
	return rootBuilder.BgYellowAnd()
}

// BgYellow returns a string where the background color is Yellow, in addition to other styles from this builder.
func (builder *Builder) BgYellow(str ...string) string {
	return builder.BgYellowAnd().applyStyle(str...)
}

// BgYellowAnd returns a Builder that generates strings where the background color is Yellow,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) BgYellowAnd() *Builder {
	if builder.bgYellow == nil {
		builder.bgYellow = createBuilder(builder, ansistyles.BgYellow.Open, ansistyles.BgYellow.Close)
	}
	return builder.bgYellow
}

// BgBlackBright returns a string where the background color is BlackBright.
func BgBlackBright(str ...string) string {
	return rootBuilder.BgBlackBright(str...)
}

// BgBlackBrightAnd returns a Builder that generates strings where the background color is BlackBright,
// and further styles can be applied via chaining.
func BgBlackBrightAnd() *Builder {
	return rootBuilder.BgBlackBrightAnd()
}

// BgBlackBright returns a string where the background color is BlackBright, in addition to other styles from this builder.
func (builder *Builder) BgBlackBright(str ...string) string {
	return builder.BgBlackBrightAnd().applyStyle(str...)
}

// BgBlackBrightAnd returns a Builder that generates strings where the background color is BlackBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) BgBlackBrightAnd() *Builder {
	if builder.bgBlackBright == nil {
		builder.bgBlackBright = createBuilder(builder, ansistyles.BgBlackBright.Open, ansistyles.BgBlackBright.Close)
	}
	return builder.bgBlackBright
}

// BgRedBright returns a string where the background color is RedBright.
func BgRedBright(str ...string) string {
	return rootBuilder.BgRedBright(str...)
}

// BgRedBrightAnd returns a Builder that generates strings where the background color is RedBright,
// and further styles can be applied via chaining.
func BgRedBrightAnd() *Builder {
	return rootBuilder.BgRedBrightAnd()
}

// BgRedBright returns a string where the background color is RedBright, in addition to other styles from this builder.
func (builder *Builder) BgRedBright(str ...string) string {
	return builder.BgRedBrightAnd().applyStyle(str...)
}

// BgRedBrightAnd returns a Builder that generates strings where the background color is RedBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) BgRedBrightAnd() *Builder {
	if builder.bgRedBright == nil {
		builder.bgRedBright = createBuilder(builder, ansistyles.BgRedBright.Open, ansistyles.BgRedBright.Close)
	}
	return builder.bgRedBright
}

// BgMagentaBright returns a string where the background color is MagentaBright.
func BgMagentaBright(str ...string) string {
	return rootBuilder.BgMagentaBright(str...)
}

// BgMagentaBrightAnd returns a Builder that generates strings where the background color is MagentaBright,
// and further styles can be applied via chaining.
func BgMagentaBrightAnd() *Builder {
	return rootBuilder.BgMagentaBrightAnd()
}

// BgMagentaBright returns a string where the background color is MagentaBright, in addition to other styles from this builder.
func (builder *Builder) BgMagentaBright(str ...string) string {
	return builder.BgMagentaBrightAnd().applyStyle(str...)
}

// BgMagentaBrightAnd returns a Builder that generates strings where the background color is MagentaBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) BgMagentaBrightAnd() *Builder {
	if builder.bgMagentaBright == nil {
		builder.bgMagentaBright = createBuilder(builder, ansistyles.BgMagentaBright.Open, ansistyles.BgMagentaBright.Close)
	}
	return builder.bgMagentaBright
}

// BgCyanBright returns a string where the background color is CyanBright.
func BgCyanBright(str ...string) string {
	return rootBuilder.BgCyanBright(str...)
}

// BgCyanBrightAnd returns a Builder that generates strings where the background color is CyanBright,
// and further styles can be applied via chaining.
func BgCyanBrightAnd() *Builder {
	return rootBuilder.BgCyanBrightAnd()
}

// BgCyanBright returns a string where the background color is CyanBright, in addition to other styles from this builder.
func (builder *Builder) BgCyanBright(str ...string) string {
	return builder.BgCyanBrightAnd().applyStyle(str...)
}

// BgCyanBrightAnd returns a Builder that generates strings where the background color is CyanBright,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) BgCyanBrightAnd() *Builder {
	if builder.bgCyanBright == nil {
		builder.bgCyanBright = createBuilder(builder, ansistyles.BgCyanBright.Open, ansistyles.BgCyanBright.Close)
	}
	return builder.bgCyanBright
}

// BgBlack returns a string where the background color is Black.
func BgBlack(str ...string) string {
	return rootBuilder.BgBlack(str...)
}

// BgBlackAnd returns a Builder that generates strings where the background color is Black,
// and further styles can be applied via chaining.
func BgBlackAnd() *Builder {
	return rootBuilder.BgBlackAnd()
}

// BgBlack returns a string where the background color is Black, in addition to other styles from this builder.
func (builder *Builder) BgBlack(str ...string) string {
	return builder.BgBlackAnd().applyStyle(str...)
}

// BgBlackAnd returns a Builder that generates strings where the background color is Black,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) BgBlackAnd() *Builder {
	if builder.bgBlack == nil {
		builder.bgBlack = createBuilder(builder, ansistyles.BgBlack.Open, ansistyles.BgBlack.Close)
	}
	return builder.bgBlack
}

// Bold returns a string with the bold modifier.
func Bold(str ...string) string {
	return rootBuilder.Bold(str...)
}

// BoldAnd returns a Builder that generates strings with the bold modifier,
// and further styles can be applied via chaining.
func BoldAnd() *Builder {
	return rootBuilder.BoldAnd()
}

// Bold returns a string with the bold modifier, in addition to other styles from this builder.
func (builder *Builder) Bold(str ...string) string {
	return builder.BoldAnd().applyStyle(str...)
}

// BoldAnd returns a Builder that generates strings with the bold modifier,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) BoldAnd() *Builder {
	if builder.bold == nil {
		builder.bold = createBuilder(builder, ansistyles.Bold.Open, ansistyles.Bold.Close)
	}
	return builder.bold
}

// Dim returns a string with the dim modifier.
func Dim(str ...string) string {
	return rootBuilder.Dim(str...)
}

// DimAnd returns a Builder that generates strings with the dim modifier,
// and further styles can be applied via chaining.
func DimAnd() *Builder {
	return rootBuilder.DimAnd()
}

// Dim returns a string with the dim modifier, in addition to other styles from this builder.
func (builder *Builder) Dim(str ...string) string {
	return builder.DimAnd().applyStyle(str...)
}

// DimAnd returns a Builder that generates strings with the dim modifier,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) DimAnd() *Builder {
	if builder.dim == nil {
		builder.dim = createBuilder(builder, ansistyles.Dim.Open, ansistyles.Dim.Close)
	}
	return builder.dim
}

// Italic returns a string with the italic modifier.
func Italic(str ...string) string {
	return rootBuilder.Italic(str...)
}

// ItalicAnd returns a Builder that generates strings with the italic modifier,
// and further styles can be applied via chaining.
func ItalicAnd() *Builder {
	return rootBuilder.ItalicAnd()
}

// Italic returns a string with the italic modifier, in addition to other styles from this builder.
func (builder *Builder) Italic(str ...string) string {
	return builder.ItalicAnd().applyStyle(str...)
}

// ItalicAnd returns a Builder that generates strings with the italic modifier,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) ItalicAnd() *Builder {
	if builder.italic == nil {
		builder.italic = createBuilder(builder, ansistyles.Italic.Open, ansistyles.Italic.Close)
	}
	return builder.italic
}

// Overline returns a string with the overline modifier.
func Overline(str ...string) string {
	return rootBuilder.Overline(str...)
}

// OverlineAnd returns a Builder that generates strings with the overline modifier,
// and further styles can be applied via chaining.
func OverlineAnd() *Builder {
	return rootBuilder.OverlineAnd()
}

// Overline returns a string with the overline modifier, in addition to other styles from this builder.
func (builder *Builder) Overline(str ...string) string {
	return builder.OverlineAnd().applyStyle(str...)
}

// OverlineAnd returns a Builder that generates strings with the overline modifier,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) OverlineAnd() *Builder {
	if builder.overline == nil {
		builder.overline = createBuilder(builder, ansistyles.Overline.Open, ansistyles.Overline.Close)
	}
	return builder.overline
}

// Inverse returns a string with the inverse modifier.
func Inverse(str ...string) string {
	return rootBuilder.Inverse(str...)
}

// InverseAnd returns a Builder that generates strings with the inverse modifier,
// and further styles can be applied via chaining.
func InverseAnd() *Builder {
	return rootBuilder.InverseAnd()
}

// Inverse returns a string with the inverse modifier, in addition to other styles from this builder.
func (builder *Builder) Inverse(str ...string) string {
	return builder.InverseAnd().applyStyle(str...)
}

// InverseAnd returns a Builder that generates strings with the inverse modifier,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) InverseAnd() *Builder {
	if builder.inverse == nil {
		builder.inverse = createBuilder(builder, ansistyles.Inverse.Open, ansistyles.Inverse.Close)
	}
	return builder.inverse
}

// Strikethrough returns a string with the strikethrough modifier.
func Strikethrough(str ...string) string {
	return rootBuilder.Strikethrough(str...)
}

// StrikethroughAnd returns a Builder that generates strings with the strikethrough modifier,
// and further styles can be applied via chaining.
func StrikethroughAnd() *Builder {
	return rootBuilder.StrikethroughAnd()
}

// Strikethrough returns a string with the strikethrough modifier, in addition to other styles from this builder.
func (builder *Builder) Strikethrough(str ...string) string {
	return builder.StrikethroughAnd().applyStyle(str...)
}

// StrikethroughAnd returns a Builder that generates strings with the strikethrough modifier,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) StrikethroughAnd() *Builder {
	if builder.strikethrough == nil {
		builder.strikethrough = createBuilder(builder, ansistyles.Strikethrough.Open, ansistyles.Strikethrough.Close)
	}
	return builder.strikethrough
}

// Reset returns a string with the reset modifier.
func Reset(str ...string) string {
	return rootBuilder.Reset(str...)
}

// ResetAnd returns a Builder that generates strings with the reset modifier,
// and further styles can be applied via chaining.
func ResetAnd() *Builder {
	return rootBuilder.ResetAnd()
}

// Reset returns a string with the reset modifier, in addition to other styles from this builder.
func (builder *Builder) Reset(str ...string) string {
	return builder.ResetAnd().applyStyle(str...)
}

// ResetAnd returns a Builder that generates strings with the reset modifier,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) ResetAnd() *Builder {
	if builder.reset == nil {
		builder.reset = createBuilder(builder, ansistyles.Reset.Open, ansistyles.Reset.Close)
	}
	return builder.reset
}

// Underline returns a string with the underline modifier.
func Underline(str ...string) string {
	return rootBuilder.Underline(str...)
}

// UnderlineAnd returns a Builder that generates strings with the underline modifier,
// and further styles can be applied via chaining.
func UnderlineAnd() *Builder {
	return rootBuilder.UnderlineAnd()
}

// Underline returns a string with the underline modifier, in addition to other styles from this builder.
func (builder *Builder) Underline(str ...string) string {
	return builder.UnderlineAnd().applyStyle(str...)
}

// UnderlineAnd returns a Builder that generates strings with the underline modifier,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) UnderlineAnd() *Builder {
	if builder.underline == nil {
		builder.underline = createBuilder(builder, ansistyles.Underline.Open, ansistyles.Underline.Close)
	}
	return builder.underline
}

// Hidden returns a string with the hidden modifier.
func Hidden(str ...string) string {
	return rootBuilder.Hidden(str...)
}

// HiddenAnd returns a Builder that generates strings with the hidden modifier,
// and further styles can be applied via chaining.
func HiddenAnd() *Builder {
	return rootBuilder.HiddenAnd()
}

// Hidden returns a string with the hidden modifier, in addition to other styles from this builder.
func (builder *Builder) Hidden(str ...string) string {
	return builder.HiddenAnd().applyStyle(str...)
}

// HiddenAnd returns a Builder that generates strings with the hidden modifier,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) HiddenAnd() *Builder {
	if builder.hidden == nil {
		builder.hidden = createBuilder(builder, ansistyles.Hidden.Open, ansistyles.Hidden.Close)
	}
	return builder.hidden
}

type generatedBuilders struct {
	blackBright     *Builder
	whiteBright     *Builder
	grey            *Builder
	bgBlue          *Builder
	bgBlueBright    *Builder
	italic          *Builder
	overline        *Builder
	red             *Builder
	bgBlack         *Builder
	bgMagenta       *Builder
	bgYellowBright  *Builder
	bgMagentaBright *Builder
	hidden          *Builder
	white           *Builder
	yellowBright    *Builder
	bgGray          *Builder
	green           *Builder
	blueBright      *Builder
	bgYellow        *Builder
	bgRed           *Builder
	bgWhite         *Builder
	bgRedBright     *Builder
	bold            *Builder
	inverse         *Builder
	strikethrough   *Builder
	bgGrey          *Builder
	reset           *Builder
	dim             *Builder
	bgCyan          *Builder
	magentaBright   *Builder
	bgBlackBright   *Builder
	yellow          *Builder
	magenta         *Builder
	greenBright     *Builder
	bgGreen         *Builder
	cyan            *Builder
	redBright       *Builder
	cyanBright      *Builder
	gray            *Builder
	bgGreenBright   *Builder
	bgCyanBright    *Builder
	bgWhiteBright   *Builder
	underline       *Builder
	black           *Builder
	blue            *Builder
}
