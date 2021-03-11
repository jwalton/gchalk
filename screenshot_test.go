package gawk

import (
	"fmt"
	"testing"
)

func TestScreenshot(t *testing.T) {
	gawk := New(ForceLevel(LevelAnsi16m))

	fmt.Println()
	fmt.Println()
	fmt.Println(`fmt.Println(gawk.Blue("This line is blue"))`)
	fmt.Println("// " + gawk.Blue("This line is blue"))
	fmt.Println()
	fmt.Println(`fmt.Println(gawk.RGB(90, 60, 245)("This line is a nice purple"))`)
	fmt.Println("// " + gawk.RGB(90, 60, 245)("This line is a nice purple"))
	fmt.Println()
	fmt.Println(`fmt.Println(gawk.WithBgBlue().BrightCyan("Cyan text with a blue background"))`)
	fmt.Println("// " + gawk.WithBgBlue().BrightCyan("Cyan text with a blue background"))
	fmt.Println()
	fmt.Println(`fmt.Println(Blue("This line is blue", Green(" and then green"), " and then blue again"))`)
	fmt.Println("// " + gawk.Blue("This line is blue", gawk.Green("and then green"), "and then blue again"))
	fmt.Println()
}
