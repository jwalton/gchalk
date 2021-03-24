package gchalk

import (
	"testing"
)

var benchGChalk = New(ForceLevel(LevelAnsi16m))

func BenchmarkColor(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		benchGChalk.Red("Hello world!")
	}
}

var red = benchGChalk.Red

func BenchmarkPreDefinedColor(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		red("Hello world!")
	}
}

func BenchmarkRgb(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		benchGChalk.RGB(10, 30, 255)("Hello world!")
	}
}

var rgbTest = benchGChalk.RGB(10, 30, 255)

func BenchmarkPreDefinedRgb(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		rgbTest("Hello world!")
	}
}
func BenchmarkComposed(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		benchGChalk.WithBgBlue().WithBold().Red("Hello world!")
	}
}

func BenchmarkNested(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		benchGChalk.Red("Hello", benchGChalk.Green("green"), "world!")
	}
}

func BenchmarkStyle(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		benchGChalk.StyleMust("red")("Hello world!")
	}
}
