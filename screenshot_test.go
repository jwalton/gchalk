package gchalk

import (
	"fmt"
	"testing"
)

func TestScreenshot(t *testing.T) {
	gchalk := New(ForceLevel(LevelAnsi16m))

	fmt.Println()
	fmt.Println()
	fmt.Println(`fmt.Println(gchalk.Blue("This line is blue"))`)
	fmt.Println("// " + gchalk.Blue("This line is blue"))
	fmt.Println()
	fmt.Println(`fmt.Println(gchalk.RGB(90, 60, 245)("This line is a nice purple"))`)
	fmt.Println("// " + gchalk.RGB(90, 60, 245)("This line is a nice purple"))
	fmt.Println()
	fmt.Println(`fmt.Println(gchalk.WithBgBlue().BrightCyan("Cyan text with a blue background"))`)
	fmt.Println("// " + gchalk.WithBgBlue().BrightCyan("Cyan text with a blue background"))
	fmt.Println()
	fmt.Println(`fmt.Println(Blue("This line is blue", Green(" and then green"), " and then blue again"))`)
	fmt.Println("// " + gchalk.Blue("This line is blue", gchalk.Green("and then green"), "and then blue again"))
	fmt.Println()
}
