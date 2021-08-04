package gchalk

import (
	"fmt"
	"strings"
	"testing"
)

func assertEqual(t *testing.T, actual string, expected string) {
	if expected != actual {
		t.Errorf("ANSI strings do not match.\nExepected: %v\nGot      : %v",
			strings.Replace(expected, "\u001B", "\\u001b", -1),
			strings.Replace(actual, "\u001B", "\\u001b", -1),
		)
	}
}

func TestStyleStrings(t *testing.T) {
	gchalk := New(ForceLevel(LevelAnsi16m))
	assertEqual(t, gchalk.Underline("foo"), "\u001B[4mfoo\u001B[24m")
	assertEqual(t, gchalk.Red("foo"), "\u001B[31mfoo\u001B[39m")
	assertEqual(t, gchalk.BgRed("foo"), "\u001B[41mfoo\u001B[49m")
}

func TestApplyingMultipleStylesAtOnce(t *testing.T) {
	gchalk := New(ForceLevel(LevelAnsi16m))
	assertEqual(t, gchalk.WithRed().WithBgGreen().Underline("foo"), "\u001B[31m\u001B[42m\u001B[4mfoo\u001B[24m\u001B[49m\u001B[39m")
	assertEqual(t, gchalk.WithUnderline().WithRed().BgGreen("foo"), "\u001B[4m\u001B[31m\u001B[42mfoo\u001B[49m\u001B[39m\u001B[24m")
}

func TestSupportNestingStyles(t *testing.T) {
	gchalk := New(ForceLevel(LevelAnsi16m))
	assertEqual(t,
		gchalk.Red("foo"+gchalk.WithUnderline().BgBlue("bar")+"!"),
		"\u001B[31mfoo\u001B[4m\u001B[44mbar\u001B[49m\u001B[24m!\u001B[39m",
	)
}

func TestSupportNestingStylesOfSameType(t *testing.T) {
	gchalk := New(ForceLevel(LevelAnsi16m))
	assertEqual(t,
		gchalk.Red("a"+gchalk.Yellow("b"+gchalk.Green("c")+"b")+"c"),
		"\u001B[31ma\u001B[33mb\u001B[32mc\u001B[33mb\u001B[31mc\u001B[39m",
	)
}

func TestNestedBoldAndDim(t *testing.T) {
	// See https://github.com/doowb/ansi-colors#nested-styling-bug, https://github.com/chalk/chalk/pull/335
	//
	// Because "dim" and "bold" both use the same closing code, lots of ANSI
	// libraries get this particular case wrong.
	gchalk := New(ForceLevel(LevelAnsi16m))

	assertEqual(t,
		gchalk.Bold("foo", gchalk.WithRed().Dim("bar"), "baz"),
		"\u001b[1mfoo \u001b[31m\u001b[2mbar\u001b[22m\u001b[1m\u001b[39m baz\u001b[22m",
	)
}

func TestResetAllStyles(t *testing.T) {
	gchalk := New(ForceLevel(LevelAnsi16m))
	assertEqual(t,
		gchalk.Reset(gchalk.WithRed().WithBgGreen().Underline("foo")+"foo"),
		"\u001B[0m\u001B[31m\u001B[42m\u001B[4mfoo\u001B[24m\u001B[49m\u001B[39mfoo\u001B[0m",
	)
}

func TestCachingMultipleStyles(t *testing.T) {
	gchalk := New(ForceLevel(LevelAnsi16m))

	red := gchalk.WithRed().Red
	green := gchalk.WithRed().Green
	redBold := gchalk.WithRed().WithRed().Bold
	greenBold := gchalk.WithRed().WithGreen().Bold

	if red("foo") == green("foo") {
		t.Errorf("red and green should produce different output")
	}

	if redBold("foo") == greenBold("foo") {
		t.Errorf("redBold and greenBold should produce different output")
	}

	if green("foo") == greenBold("foo") {
		t.Errorf("green and greenBold should produce different output")
	}
}

func TestAliasGrayToGrey(t *testing.T) {
	gchalk := New(ForceLevel(LevelAnsi16m))

	assertEqual(t, gchalk.Gray("foo"), "\u001B[90mfoo\u001B[39m")
	assertEqual(t, gchalk.Grey("foo"), "\u001B[90mfoo\u001B[39m")
}

func TestSupportVariableNumberOfArguments(t *testing.T) {
	gchalk := New(ForceLevel(LevelAnsi16m))
	assertEqual(t, gchalk.Red("foo", "bar"), "\u001B[31mfoo bar\u001B[39m")
}

func TestNoEscapeCodesIfInputIsEmpty(t *testing.T) {
	gchalk := New(ForceLevel(LevelAnsi16m))
	assertEqual(t, gchalk.Red(), "")
	assertEqual(t, gchalk.Red(""), "")
	assertEqual(t, gchalk.WithRed().WithBlue().Black(), "")
}

func TestLineBreaksShouldOpenAndCloseColors(t *testing.T) {
	gchalk := New(ForceLevel(LevelAnsi16m))
	assertEqual(t,
		gchalk.Grey("hello\nworld"),
		"\u001B[90mhello\u001B[39m\n\u001B[90mworld\u001B[39m",
	)
}

func TestLineBreaksShouldOpenAndCLoseColorsWithCRLF(t *testing.T) {
	gchalk := New(ForceLevel(LevelAnsi16m))
	assertEqual(t,
		gchalk.Grey("hello\r\nworld"),
		"\u001B[90mhello\u001B[39m\r\n\u001B[90mworld\u001B[39m",
	)
}

func TestSupportsBrightBlackColor(t *testing.T) {
	gchalk := New(ForceLevel(LevelAnsi16m))
	assertEqual(t,
		gchalk.BrightBlack("foo"),
		"\u001B[90mfoo\u001B[39m",
	)
}

func TestThemes(t *testing.T) {
	gchalk := New(ForceLevel(LevelAnsi16m))

	error := gchalk.WithBold().Red
	result := error("Test")
	assertEqual(t, result, "\u001b[1m\u001b[31mTest\u001b[39m\u001b[22m")
}

func TestSprintfSubstitution(t *testing.T) {
	gchalk := New(ForceLevel(LevelAnsi16m))

	assertEqual(t,
		fmt.Sprintf(gchalk.Green("Hello %s"), "Jason"),
		"\u001b[32mHello Jason\u001b[39m",
	)
}

func TestColorDowngradeNone(t *testing.T) {
	gchalkBW := New(ForceLevel(LevelNone))

	assertEqual(t,
		gchalkBW.Green("Hello"),
		"Hello",
	)

	assertEqual(t,
		gchalkBW.Ansi256(201)("Hello256"),
		"Hello256",
	)

	assertEqual(t,
		gchalkBW.Hex("#0C0")("HelloHex"),
		"HelloHex",
	)
}

func TestColorDowngradeBasic(t *testing.T) {
	gchalkBasic := New(ForceLevel(LevelBasic))
	//gchalk256 := New(ForceLevel(LevelAnsi256))

	assertEqual(t,
		gchalkBasic.Green("Hello"),
		"\u001b[32mHello\u001b[39m",
	)

	assertEqual(t,
		gchalkBasic.Ansi256(40)("Hello256"),
		"\u001b[32mHello256\u001b[39m",
	)

	assertEqual(t,
		gchalkBasic.Hex("#0C0")("HelloHex"),
		"\u001b[32mHelloHex\u001b[39m",
	)
}

func TestColorDowngradeAnsi256(t *testing.T) {
	gchalk256 := New(ForceLevel(LevelAnsi256))

	assertEqual(t,
		gchalk256.Green("Hello"),
		"\u001b[32mHello\u001b[39m",
	)

	assertEqual(t,
		gchalk256.Ansi256(40)("Hello256"),
		"\u001b[38;5;40mHello256\u001b[39m",
	)

	assertEqual(t,
		gchalk256.Hex("#0C0")("HelloHex"),
		"\u001b[38;5;40mHelloHex\u001b[39m",
	)
}

func TestStyle(t *testing.T) {
	gchalk := New(ForceLevel(LevelAnsi16m))

	assertEqual(t, gchalk.StyleMust("underline")("foo"), "\u001B[4mfoo\u001B[24m")
	assertEqual(t, gchalk.StyleMust("red")("foo"), "\u001B[31mfoo\u001B[39m")
	assertEqual(t, gchalk.StyleMust("brightRed")("foo"), "\u001B[91mfoo\u001B[39m")
	assertEqual(t, gchalk.StyleMust("bgRed")("foo"), "\u001B[41mfoo\u001B[49m")
	// Should be case insensitive.
	assertEqual(t, gchalk.StyleMust("bgred")("foo"), "\u001B[41mfoo\u001B[49m")

	assertEqual(t,
		gchalk.StyleMust("red", "bgGreen", "underline")("foo"),
		"\u001B[31m\u001B[42m\u001B[4mfoo\u001B[24m\u001B[49m\u001B[39m",
	)

	styler, err := gchalk.Style("idonotexist")
	str := styler("foo")
	assertEqual(t, str, "foo")
	assertEqual(t, fmt.Sprintf("%v", err), "No such style: idonotexist")
}

func TestStyleWithHex(t *testing.T) {
	gchalk := New(ForceLevel(LevelAnsi16m))

	assertEqual(t, gchalk.StyleMust("#FF00FF")("foo"), "\u001b[38;2;255;0;255mfoo\u001b[39m")
	assertEqual(t, gchalk.StyleMust("bg#FF00FF")("foo"), "\u001b[48;2;255;0;255mfoo\u001b[49m")
}

func TestPaint(t *testing.T) {
	gchalk := New(ForceLevel(LevelAnsi16m))

	assertEqual(t, gchalk.WithRed().Paint("foo"), "\u001B[31mfoo\u001B[39m")
}

func TestSprintf(t *testing.T) {
	gchalk := New(ForceLevel(LevelAnsi16m))

	assertEqual(t, gchalk.WithRed().Sprintf("Hello %s", "World"), "\u001B[31mHello World\u001B[39m")
}
