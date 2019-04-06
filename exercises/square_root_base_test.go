package exercises

import (
	"testing"
	"unsafe"
)

const (
	FLOAT_NUMBER float64 = 1 << 20
	FLOAT_SIZE   int64   = int64(unsafe.Sizeof(float64(1.0)))
)

func BenchmarkSquareRootBaseEstimating(b *testing.B) {
	b.SetBytes(FLOAT_SIZE)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			squareRootBaseEstimating(FLOAT_NUMBER)
		}
	})
}

func BenchmarkSquareRootBaseEstimatingFormula(b *testing.B) {
	b.SetBytes(FLOAT_SIZE)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			squareRootBaseEstimatingFormula(FLOAT_NUMBER)
		}
	})
}

func BenchmarkSquareRootBaseEstimatingBuiltin(b *testing.B) {
	b.SetBytes(FLOAT_SIZE)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			squareRootBaseEstimatingBuiltin(FLOAT_NUMBER)
		}
	})
}
