package gawk

import (
	"testing"
)

var benchGawk = New(ForceLevel(LevelAnsi16m))

func BenchmarkColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchGawk.Red("Hello world!")
	}
}

func BenchmarkRgb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchGawk.RGB(10, 30, 255)("Hello world!")
	}
}

func BenchmarkColorNoSetup(b *testing.B) {
	benchGawk.Red("Hello world!")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchGawk.Red("Hello world!")
	}
}

func BenchmarkComposed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchGawk.WithBgBlue().WithBold().Red("Hello world!")
	}
}

func BenchmarkNested(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchGawk.Red("Hello", benchGawk.Green("green"), "world!")
	}
}
