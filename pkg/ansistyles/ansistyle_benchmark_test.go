package ansistyles

import (
	"strings"
	"testing"
)

func BenchmarkAnsi(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Ansi(7)
	}
}

func BenchmarkAnsi16m(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Ansi16m(10, 30, 255)
	}
}

func BenchmarkWriteStringAnsi(b *testing.B) {
	out := strings.Builder{}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		WriteStringAnsi(&out, 7)
	}
}
