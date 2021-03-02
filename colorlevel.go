package gawk

//go:generate stringer -type=ColorLevel

// ColorLevel represents the ANSI color level supported by the terminal.
type ColorLevel int

const (
	// LevelNone represents a terminal that does not support color at all.
	LevelNone ColorLevel = 0
	// LevelBasic represents a terminal with basic 16 color support.
	LevelBasic ColorLevel = 1
	// LevelAnsi256 represents a terminal with 256 color support.
	LevelAnsi256 ColorLevel = 2
	// LevelAnsi16m represents a terminal with full true color support.
	LevelAnsi16m ColorLevel = 3
)
