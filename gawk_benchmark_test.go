package gawk

import (
	"testing"
)

func BenchmarkColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Red("Hello world!")
	}
}

func BenchmarkComposed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BgBlueAnd().BoldAnd().Red("Hello world!")
	}
}

func BenchmarkNested(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Red("Hello", Green("green"), "world!")
	}
}
