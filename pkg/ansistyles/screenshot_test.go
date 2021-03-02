package ansistyles

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"testing"
)

func printStyles(styles *map[string]CSPair) {
	bgRegex := regexp.MustCompile("^bg[^B]")
	const width = 55
	lineLength := 0

	keys := make([]string, len(*styles))
	i := 0
	for key := range *styles {
		keys[i] = key
		i = i + 1
	}
	sort.Strings(keys)

	for _, key := range keys {
		value := (*styles)[key]
		code := value.Open
		projectedLength := lineLength + len(key) + 1

		if key == "reset" || key == "hidden" || key == "grey" || strings.Contains(key, "Bright") {
			continue
		}

		if bgRegex.MatchString(key) {
			code = Black.Open + code
		}

		if width < projectedLength {
			fmt.Print("\n")
			return
		}

		fmt.Print(code + key + Reset.Close + " ")
		lineLength = projectedLength
	}
	fmt.Print("\n")
}

func TestScreenshot(t *testing.T) {
	printStyles(&Modifier)
	printStyles(&Color)
	printStyles(&BgColor)
}
