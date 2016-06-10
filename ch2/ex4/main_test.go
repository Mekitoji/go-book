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
