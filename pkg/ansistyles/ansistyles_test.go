package ansistyles

import (
	"strings"
	"testing"
)

func assertEqual(t *testing.T, actual string, expected string) {
	if expected != actual {
		t.Errorf("Exepected: %v got: %v",
			strings.Replace(expected, "\u001B", "\\u001b", -1),
			strings.Replace(actual, "\u001B", "\\u001b", -1),
		)
	}
}

func assertEqualUint8(t *testing.T, actual uint8, expected uint8) {
	if expected != actual {
		t.Errorf("Exepected: %v got: %v", expected, actual)
	}
}

func TestANSIEscapeCodes(t *testing.T) {
	assertEqual(t, Green.Open, "\u001B[32m")
	assertEqual(t, BgGreen.Open, "\u001B[42m")
	assertEqual(t, Green.Close, "\u001B[39m")
	assertEqual(t, Gray.Open, Grey.Open)
}

func TestGroupRelatedCodesIntoCategories(t *testing.T) {
	assertEqual(t, Color["magenta"].Open, Magenta.Open)
	assertEqual(t, BgColor["bgYellow"].Open, BgYellow.Open)
	assertEqual(t, Modifier["bold"].Open, Bold.Open)
}

func TestColorANSIEscapeCodes(t *testing.T) {
	assertEqual(t, Styles["green"].Open, "\u001B[32m")
	assertEqual(t, Styles["bgGreen"].Open, "\u001B[42m")
	assertEqual(t, Styles["green"].Close, "\u001B[39m")
	assertEqual(t, Styles["grey"].Open, Color["gray"].Open)
}

func TestHexToRGB(t *testing.T) {
	r, g, b := HexToRGB("#102132")
	if r != 16 || g != 33 || b != 50 {
		t.Errorf("#102132 - Expected: 16, 33, 50, got: %v, %v, %v", r, g, b)
	}

	r, g, b = HexToRGB("#ff00ff")
	if r != 255 || g != 0 || b != 255 {
		t.Errorf("#ff00ff - Expected: 255, 0, 255, got: %v, %v, %v", r, g, b)
	}

	r, g, b = HexToRGB("#FF00FF")
	if r != 255 || g != 0 || b != 255 {
		t.Errorf("#FF00FF - Expected: 255, 0, 255, got: %v, %v, %v", r, g, b)
	}

	r, g, b = HexToRGB("ff00ff")
	if r != 255 || g != 0 || b != 255 {
		t.Errorf("ff00ff - Expected: 255, 0, 255, got: %v, %v, %v", r, g, b)
	}
}

func TestConversionToAnsiColors(t *testing.T) {
	assertEqual(t, Ansi(RGBToAnsi(255, 255, 255)), "\u001B[97m")
	assertEqual(t, Ansi(RGBToAnsi(HexToRGB("#990099"))), "\u001B[35m")
	assertEqual(t, Ansi(RGBToAnsi(HexToRGB("#FF00FF"))), "\u001B[35m")
	assertEqual(t, Ansi(Ansi256ToAnsi(201)), "\u001B[35m")

	assertEqual(t, BgAnsi(RGBToAnsi(255, 255, 255)), "\u001B[107m")
	assertEqual(t, BgAnsi(RGBToAnsi(HexToRGB("#990099"))), "\u001B[45m")
	assertEqual(t, BgAnsi(RGBToAnsi(HexToRGB("#FF00FF"))), "\u001B[45m")
	assertEqual(t, BgAnsi(Ansi256ToAnsi(201)), "\u001B[45m")
}

func TestConversionToAnsi256Colors(t *testing.T) {
	assertEqual(t, Ansi256(RGBToAnsi256(255, 255, 255)), "\u001B[38;5;231m")
	assertEqual(t, Ansi256(HexToAnsi256("#990099")), "\u001B[38;5;127m")
	assertEqual(t, Ansi256(HexToAnsi256("#FF00FF")), "\u001B[38;5;201m")

	assertEqual(t, BgAnsi256(RGBToAnsi256(255, 255, 255)), "\u001B[48;5;231m")
	assertEqual(t, BgAnsi256(HexToAnsi256("#990099")), "\u001B[48;5;127m")
	assertEqual(t, BgAnsi256(HexToAnsi256("#FF00FF")), "\u001B[48;5;201m")
}

func TestConversionToAnsi16MillionColors(t *testing.T) {
	assertEqual(t, Ansi16m(255, 255, 255), "\u001B[38;2;255;255;255m")
	assertEqual(t, Ansi16m(HexToRGB("#990099")), "\u001B[38;2;153;0;153m")
	assertEqual(t, Ansi16m(HexToRGB("#FF00FF")), "\u001B[38;2;255;0;255m")

	assertEqual(t, BgAnsi16m(255, 255, 255), "\u001B[48;2;255;255;255m")
	assertEqual(t, BgAnsi16m(HexToRGB("#990099")), "\u001B[48;2;153;0;153m")
	assertEqual(t, BgAnsi16m(HexToRGB("#FF00FF")), "\u001B[48;2;255;0;255m")
}

func TestExportRawAnsiEscapeCodes(t *testing.T) {
	assertEqualUint8(t, CloseCode(0), 0)
	assertEqualUint8(t, CloseCode(1), 22)
	assertEqualUint8(t, CloseCode(91), 39)
	assertEqualUint8(t, CloseCode(40), 49)
	assertEqualUint8(t, CloseCode(100), 49)
}

func TestAnsi16m(t *testing.T) {
	assertEqual(t, Ansi16m(123, 45, 67), "\u001B[38;2;123;45;67m")
}
