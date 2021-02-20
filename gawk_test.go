package gawk

import (
	"fmt"
	"strings"
	"testing"
)

// TODO: Set level in tests, so they pass on systems where ANSI is not supported.

func assertEqual(t *testing.T, actual string, expected string) {
	if expected != actual {
		t.Errorf("ANSI strings do not match.\nExepected: %v\nGot      : %v",
			strings.Replace(expected, "\u001B", "\\u001b", -1),
			strings.Replace(actual, "\u001B", "\\u001b", -1),
		)
	}
}

func TestThemes(t *testing.T) {
	error := BoldAnd().Red
	result := error("Test")
	assertEqual(t, result, "\u001b[1m\u001b[31mTest\u001b[39m\u001b[22m")
}

func TestSprintfSubstitution(t *testing.T) {
	assertEqual(t,
		fmt.Sprintf(Green("Hello %s"), "Jason"),
		"\u001b[32mHello Jason\u001b[39m",
	)
}
