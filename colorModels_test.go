package gchalk

import (
	"testing"
)

func TestAnsi(t *testing.T) {
	gchalk := New(ForceLevel(LevelAnsi16m))
	assertEqual(t, gchalk.Ansi(31)("foo"), "\u001b[31mfoo\u001b[39m")
	assertEqual(t, gchalk.BgAnsi(31)("foo"), "\u001b[41mfoo\u001b[49m")
	assertEqual(t, gchalk.Ansi(1)("foo"), "\u001b[1mfoo\u001b[22m")

	SetLevel(LevelAnsi16m)
	assertEqual(t, Ansi(31)("foo"), "\u001b[31mfoo\u001b[39m")
	assertEqual(t, BgAnsi(31)("foo"), "\u001b[41mfoo\u001b[49m")
	assertEqual(t, Ansi(1)("foo"), "\u001b[1mfoo\u001b[22m")
}

func TestAnsi256(t *testing.T) {
	gchalk := New(ForceLevel(LevelAnsi16m))
	assertEqual(t, gchalk.Ansi256(201)("foo"), "\u001b[38;5;201mfoo\u001b[39m")
	assertEqual(t, gchalk.BgAnsi256(201)("foo"), "\u001b[48;5;201mfoo\u001b[49m")

	SetLevel(LevelAnsi16m)
	assertEqual(t, Ansi256(201)("foo"), "\u001b[38;5;201mfoo\u001b[39m")
	assertEqual(t, BgAnsi256(201)("foo"), "\u001b[48;5;201mfoo\u001b[49m")
}

func TestRGB(t *testing.T) {
	gchalk := New(ForceLevel(LevelAnsi16m))
	assertEqual(t, gchalk.RGB(100, 200, 255)("foo"), "\u001b[38;2;100;200;255mfoo\u001b[39m")
	assertEqual(t, gchalk.BgRGB(100, 200, 255)("foo"), "\u001b[48;2;100;200;255mfoo\u001b[49m")

	SetLevel(LevelAnsi16m)
	assertEqual(t, RGB(100, 200, 255)("foo"), "\u001b[38;2;100;200;255mfoo\u001b[39m")
	assertEqual(t, BgRGB(100, 200, 255)("foo"), "\u001b[48;2;100;200;255mfoo\u001b[49m")
}

func TestHex(t *testing.T) {
	gchalk := New(ForceLevel(LevelAnsi16m))
	assertEqual(t, gchalk.Hex("#64c8ff")("foo"), "\u001b[38;2;100;200;255mfoo\u001b[39m")
	assertEqual(t, gchalk.BgHex("#64c8ff")("foo"), "\u001b[48;2;100;200;255mfoo\u001b[49m")

	SetLevel(LevelAnsi16m)
	assertEqual(t, Hex("#64c8ff")("foo"), "\u001b[38;2;100;200;255mfoo\u001b[39m")
	assertEqual(t, BgHex("#64c8ff")("foo"), "\u001b[48;2;100;200;255mfoo\u001b[49m")
}
