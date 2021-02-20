package gawk

import (
	"fmt"
	"strings"

	"github.com/jwalton/go-ansistyles"
)

type stylerData struct {
	open     string
	close    string
	openAll  string
	closeAll string
	parent   *stylerData
}

// Builder is an intermediate result used to chain together styles.
type Builder struct {
	generatedBuilders
	styler      *stylerData
	rootBuilder *Builder
	Level       int
}

// Instance creates a new instance of Gawk.
func Instance() *Builder {
	builder := &Builder{styler: nil}
	builder.rootBuilder = builder
	// FIXME: Auto-detect level
	// FIXME: Instead of storing level on the builder, store it in a config object that's shared across builders.
	// FIXME: Add option to `Instance()` for setting level.
	// FIXME: Look into storing closures in Builder, see if this speeds things up, since we can get rid of an if.
	// FIXME: Export an interface instead of exporting Builder directly.
	builder.Level = 3
	return builder
}

var rootBuilder = Instance()

func createBuilder(builder *Builder, open string, close string) *Builder {
	var parent *stylerData
	if builder.styler != nil {
		parent = builder.styler
	}

	openAll := open
	closeAll := close
	if parent != nil {
		openAll = parent.openAll + open
		closeAll = close + parent.closeAll
	}

	return &Builder{
		styler: &stylerData{
			open:     open,
			close:    close,
			openAll:  openAll,
			closeAll: closeAll,
			parent:   parent,
		},
	}
}

func (builder *Builder) applyStyle(strs ...string) string {
	if /* FIXME: builder.level <= 0 || */ len(strs) == 0 {
		return ""
	}

	str := strings.Join(strs, " ")
	if str == "" {
		return ""
	}

	styler := builder.styler
	openAll := styler.openAll
	closeAll := styler.closeAll

	if styler == nil {
		return str
	}

	if strings.Contains(str, "\u001B") {
		for styler != nil {
			// Replace any instances already present with a re-opening code
			// otherwise only the part of the string until said closing code
			// will be colored, and the rest will simply be 'plain'.
			str = strings.ReplaceAll(str, styler.close, styler.open)

			styler = styler.parent
		}
	}

	// We can move both next actions out of loop, because remaining actions in loop won't have
	// any/visible effect on parts we add here. Close the styling before a linebreak and reopen
	// after next line to fix a bleed issue on macOS: https://github.com/chalk/chalk/pull/92
	if strings.Contains(str, "\n") {
		str = stringEncaseCRLF(str, closeAll, openAll)
	}

	return fmt.Sprintf("%s%s%s", openAll, str, closeAll)
}

// Black returns a string where the color is black.
func RGB(r uint8, g uint8, b uint8) func(strs ...string) string {
	return rootBuilder.RGB(r, g, b)
}

// BlackAnd returns a Builder that generates strings where the color is black,
// and further styles can be applied via chaining.
func RGBAnd(r uint8, g uint8, b uint8) *Builder {
	return rootBuilder.RGBAnd(r, g, b)
}

// Black returns a string where the color is black, in addition to other styles from this builder.
func (builder *Builder) RGB(r uint8, g uint8, b uint8) func(strs ...string) string {
	return builder.RGBAnd(r, g, b).applyStyle
}

// BlackAnd returns a Builder that generates strings where the color is black,
// in addition to other styles from this builder, and further styles can be applied via chaining.
func (builder *Builder) RGBAnd(r uint8, g uint8, b uint8) *Builder {
	if builder.black == nil {
		builder.black = createBuilder(builder, ansistyles.Ansi16m(r, g, b), ansistyles.Close)
	}
	return builder.black
}
