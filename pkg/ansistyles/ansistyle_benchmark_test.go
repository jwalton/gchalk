package ansistyles

import (
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
