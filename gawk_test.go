package gawk

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
	gawk := New(ForceLevel(LevelAnsi16m))
	assertEqual(t, gawk.Underline("foo"), "\u001B[4mfoo\u001B[24m")
	assertEqual(t, gawk.Red("foo"), "\u001B[31mfoo\u001B[39m")
	assertEqual(t, gawk.BgRed("foo"), "\u001B[41mfoo\u001B[49m")
}

func TestApplyingMultipleStylesAtOnce(t *testing.T) {
	gawk := New(ForceLevel(LevelAnsi16m))
	assertEqual(t, gawk.WithRed().WithBgGreen().Underline("foo"), "\u001B[31m\u001B[42m\u001B[4mfoo\u001B[24m\u001B[49m\u001B[39m")
	assertEqual(t, gawk.WithUnderline().WithRed().BgGreen("foo"), "\u001B[4m\u001B[31m\u001B[42mfoo\u001B[49m\u001B[39m\u001B[24m")
}

func TestSupportNestingStyles(t *testing.T) {
	gawk := New(ForceLevel(LevelAnsi16m))
	assertEqual(t,
		gawk.Red("foo"+gawk.WithUnderline().BgBlue("bar")+"!"),
		"\u001B[31mfoo\u001B[4m\u001B[44mbar\u001B[49m\u001B[24m!\u001B[39m",
	)
}

func TestSupportNestingStylesOfSameType(t *testing.T) {
	gawk := New(ForceLevel(LevelAnsi16m))
	assertEqual(t,
		gawk.Red("a"+gawk.Yellow("b"+gawk.Green("c")+"b")+"c"),
		"\u001B[31ma\u001B[33mb\u001B[32mc\u001B[33mb\u001B[31mc\u001B[39m",
	)
}

func TestNestedBoldAndDim(t *testing.T) {
	// See https://github.com/doowb/ansi-colors#nested-styling-bug, https://github.com/chalk/chalk/pull/335
	//
	// Because "dim" and "bold" both use the same closing code, lots of ANSI
	// libraries get this particular case wrong.
	gawk := New(ForceLevel(LevelAnsi16m))

	assertEqual(t,
		gawk.Bold("foo", gawk.WithRed().Dim("bar"), "baz"),
		"\u001b[1mfoo \u001b[31m\u001b[2mbar\u001b[22m\u001b[1m\u001b[39m baz\u001b[22m",
	)
}

func TestResetAllStyles(t *testing.T) {
	gawk := New(ForceLevel(LevelAnsi16m))
	assertEqual(t,
		gawk.Reset(gawk.WithRed().WithBgGreen().Underline("foo")+"foo"),
		"\u001B[0m\u001B[31m\u001B[42m\u001B[4mfoo\u001B[24m\u001B[49m\u001B[39mfoo\u001B[0m",
	)
}

func TestCachingMultipleStyles(t *testing.T) {
	gawk := New(ForceLevel(LevelAnsi16m))

	red := gawk.WithRed().Red
	green := gawk.WithRed().Green
	redBold := gawk.WithRed().WithRed().Bold
	greenBold := gawk.WithRed().WithGreen().Bold

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
	gawk := New(ForceLevel(LevelAnsi16m))

	assertEqual(t, gawk.Gray("foo"), "\u001B[90mfoo\u001B[39m")
	assertEqual(t, gawk.Grey("foo"), "\u001B[90mfoo\u001B[39m")
}

func TestSupportVariableNumberOfArguments(t *testing.T) {
	gawk := New(ForceLevel(LevelAnsi16m))
	assertEqual(t, gawk.Red("foo", "bar"), "\u001B[31mfoo bar\u001B[39m")
}

func TestNoEscapeCodesIfInputIsEmpty(t *testing.T) {
	gawk := New(ForceLevel(LevelAnsi16m))
	assertEqual(t, gawk.Red(), "")
	assertEqual(t, gawk.Red(""), "")
	assertEqual(t, gawk.WithRed().WithBlue().Black(), "")
}

func TestLineBreaksShouldOpenAndCloseColors(t *testing.T) {
	gawk := New(ForceLevel(LevelAnsi16m))
	assertEqual(t,
		gawk.Grey("hello\nworld"),
		"\u001B[90mhello\u001B[39m\n\u001B[90mworld\u001B[39m",
	)
}

func TestLineBreaksShouldOpenAndCLoseColorsWithCRLF(t *testing.T) {
	gawk := New(ForceLevel(LevelAnsi16m))
	assertEqual(t,
		gawk.Grey("hello\r\nworld"),
		"\u001B[90mhello\u001B[39m\r\n\u001B[90mworld\u001B[39m",
	)
}

func TestSupportsBrightBlackColor(t *testing.T) {
	gawk := New(ForceLevel(LevelAnsi16m))
	assertEqual(t,
		gawk.BrightBlack("foo"),
		"\u001B[90mfoo\u001B[39m",
	)
}

func TestThemes(t *testing.T) {
	gawk := New(ForceLevel(LevelAnsi16m))

	error := gawk.WithBold().Red
	result := error("Test")
	assertEqual(t, result, "\u001b[1m\u001b[31mTest\u001b[39m\u001b[22m")
}

func TestSprintfSubstitution(t *testing.T) {
	gawk := New(ForceLevel(LevelAnsi16m))

	assertEqual(t,
		fmt.Sprintf(gawk.Green("Hello %s"), "Jason"),
		"\u001b[32mHello Jason\u001b[39m",
	)
}

func TestColorDowngradeNone(t *testing.T) {
	gawkBW := New(ForceLevel(LevelNone))

	assertEqual(t,
		gawkBW.Green("Hello"),
		"Hello",
	)

	assertEqual(t,
		gawkBW.Ansi256(201)("Hello256"),
		"Hello256",
	)

	assertEqual(t,
		gawkBW.Hex("#0C0")("HelloHex"),
		"HelloHex",
	)
}

func TestColorDowngradeBasic(t *testing.T) {
	gawkBasic := New(ForceLevel(LevelBasic))
	//gawk256 := New(ForceLevel(LevelAnsi256))

	assertEqual(t,
		gawkBasic.Green("Hello"),
		"\u001b[32mHello\u001b[39m",
	)

	assertEqual(t,
		gawkBasic.Ansi256(40)("Hello256"),
		"\u001b[32mHello256\u001b[39m",
	)

	assertEqual(t,
		gawkBasic.Hex("#0C0")("HelloHex"),
		"\u001b[32mHelloHex\u001b[39m",
	)
}

func TestColorDowngradeAnsi256(t *testing.T) {
	gawk256 := New(ForceLevel(LevelAnsi256))

	assertEqual(t,
		gawk256.Green("Hello"),
		"\u001b[32mHello\u001b[39m",
	)

	assertEqual(t,
		gawk256.Ansi256(40)("Hello256"),
		"\u001b[38;5;40mHello256\u001b[39m",
	)

	assertEqual(t,
		gawk256.Hex("#0C0")("HelloHex"),
		"\u001b[38;5;40mHelloHex\u001b[39m",
	)
}

func TestStyle(t *testing.T) {
	gawk := New(ForceLevel(LevelAnsi16m))

	assertEqual(t, gawk.StyleMust("underline")("foo"), "\u001B[4mfoo\u001B[24m")
	assertEqual(t, gawk.StyleMust("red")("foo"), "\u001B[31mfoo\u001B[39m")
	assertEqual(t, gawk.StyleMust("brightRed")("foo"), "\u001B[91mfoo\u001B[39m")
	assertEqual(t, gawk.StyleMust("bgRed")("foo"), "\u001B[41mfoo\u001B[49m")

	assertEqual(t,
		gawk.StyleMust("red", "bgGreen", "underline")("foo"),
		"\u001B[31m\u001B[42m\u001B[4mfoo\u001B[24m\u001B[49m\u001B[39m",
	)

	styler, err := gawk.Style("idonotexist")
	str := styler("foo")
	assertEqual(t, str, "foo")
	assertEqual(t, fmt.Sprintf("%v", err), "No such style: idonotexist")
}
