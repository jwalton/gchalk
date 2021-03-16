package gchalk

import (
	"testing"
)

var benchGChalk = New(ForceLevel(LevelAnsi16m))

func BenchmarkColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchGChalk.Red("Hello world!")
	}
}

func BenchmarkRgb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchGChalk.RGB(10, 30, 255)("Hello world!")
	}
}

func BenchmarkColorNoSetup(b *testing.B) {
	benchGChalk.Red("Hello world!")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchGChalk.Red("Hello world!")
	}
}

func BenchmarkComposed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchGChalk.WithBgBlue().WithBold().Red("Hello world!")
	}
}

func BenchmarkNested(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchGChalk.Red("Hello", benchGChalk.Green("green"), "world!")
	}
}

func BenchmarkStyle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchGChalk.StyleMust("red")("Hello world!")
	}
}
