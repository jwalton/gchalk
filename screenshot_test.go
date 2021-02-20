package gawk

import (
	"fmt"
	"testing"
)

func TestScreenshot(t *testing.T) {
	fmt.Println()
	fmt.Println()
	fmt.Println(`fmt.Println(Blue("This line is blue"))`)
	fmt.Println("// " + Blue("This line is blue"))
	fmt.Println()
	fmt.Println(`fmt.Println(RGB(90, 60, 245)("This line is a nice purple"))`)
	fmt.Println("// " + RGB(90, 60, 245)("This line is a nice purple"))
	fmt.Println()
	fmt.Println(`fmt.Println(CyanBrightAnd().BgBlue("Cyan text with a blue background"))`)
	fmt.Println("// " + CyanBrightAnd().BgBlue("Cyan text with a blue background"))
	fmt.Println()
	fmt.Println(`fmt.Println(Blue("This line is blue", Green(" and then green"), " and then blue again"))`)
	fmt.Println("// " + Blue("This line is blue", Green("and then green"), "and then blue again"))
	fmt.Println()
}
