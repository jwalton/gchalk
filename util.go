package gchalk

import (
	"regexp"
)

var crlfRegex = regexp.MustCompile(`(\r\n|\n)`)

func stringEncaseCRLF(str string, prefix string, postfix string) string {
	return crlfRegex.ReplaceAllString(str, prefix+"${1}"+postfix)
}
