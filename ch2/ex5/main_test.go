package popcount

import "testing"

const test = 0x123ADCF42397

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(test)
	}
}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop(test)
	}
}

func BenchmarkShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Shift(test)
	}
}

func BenchmarkClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Clear(test)
	}
}

// BenchmarkPopCount-8     300000000                5.88 ns/op
// BenchmarkLoop-8         100000000               12.2 ns/op
// BenchmarkShift-8        20000000                97.2 ns/op
// BenchmarkClear-8        100000000               20.9 ns/op
// ok      github.com/mekitoji/go-book/ch2/ex5     7.897s
