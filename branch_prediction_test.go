package main

import (
	"testing"
	"unsafe"
)

var (
	ELEMENTS_NUM int = 1 << 20                    // 2^20 == 1048576 elements (int64 type)
	ELEMENT_SIZE int = int(unsafe.Sizeof(int(1))) // one int represented by, e.g. 8 bytes
)

// 396 ns/op
func BenchmarkBranchPrediction(b *testing.B) {
	slice := initSlice(ELEMENTS_NUM)
	b.SetBytes(int64(len(slice) * ELEMENT_SIZE))

	b.ResetTimer() // resetting bench timer to exclude slice allocation time

	for i := 1; i < b.N; i++ {
		branchPredictionIfElse(slice)
	}
}

// optimized loop, without if-else control flow (replaced by bitwise ops)
// 77.5 ns/op
// approximate x5 times faster than if-else control flow
func BenchmarkBranchPredictionBitwise(b *testing.B) {
	slice := initSlice(ELEMENTS_NUM)
	b.SetBytes(int64(len(slice) * ELEMENT_SIZE))

	b.ResetTimer()

	for i := 1; i < b.N; i++ {
		branchPredictionBitwise(slice)
	}
}

// -cpu 4
func BenchmarkParallelBranchPredictionIfElse(b *testing.B) {
	slice := initSlice(ELEMENTS_NUM)
	b.SetBytes(int64(len(slice) * ELEMENT_SIZE))

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			branchPredictionIfElse(slice)
		}
	})
}

func BenchmarkParallelBranchPredictionBitwise(b *testing.B) {
	slice := initSlice(ELEMENTS_NUM)
	b.SetBytes(int64(len(slice) * ELEMENT_SIZE))

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			branchPredictionBitwise(slice)
		}
	})
}

// (OLD! but valid.)
// $ go test -bench BranchPrediction

// n == 1, sum == 0goos: linux
// goarch: amd64
// pkg: github.com/itnelo/go-training
// BenchmarkBranchPrediction-8
// n == 100, sum == 9879
// n == 10000, sum == 951069
// n == 1000000, sum == 95863067
// n == 5000000, sum == 478529631 5000000	       396 ns/op

// n == 1, sum == 0BenchmarkBranchPredictionBitwise-8
// n == 100, sum == 9879
// n == 10000, sum == 951069
// n == 1000000, sum == 95863067
// n == 20000000, sum == 191543476720000000	        77.5 ns/op
// PASS
// ok  	github.com/itnelo/go-training	4.666s
