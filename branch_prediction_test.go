package main

import (
	"fmt"
	"testing"
)

// 396 ns/op
func BenchmarkBranchPrediction(b *testing.B) {
	slice := initSlice(b.N)
	b.ResetTimer() // resetting bench timer to exclude slice allocation time

	var sum int
	for i := 1; i < 100; i++ {
		sum = branchPredictionIfElse(slice)
	}

	fmt.Printf("\nn == %v, sum == %v", b.N, sum)
}

// optimized loop, without if-else control flow (replaced by bitwise ops)
// 77.5 ns/op
// approximate x5 times faster than if-else control flow
func BenchmarkBranchPredictionBitwise(b *testing.B) {
	slice := initSlice(b.N)
	b.ResetTimer()

	var sum int
	for i := 1; i < 100; i++ {
		sum = branchPredictionBitwise(slice)
	}

	fmt.Printf("\nn == %v, sum == %v", b.N, sum)
}

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
