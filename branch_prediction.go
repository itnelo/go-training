package main

import (
	"fmt"
	"math/rand"
)

func initSlice(n int) (slice []int) {
	rand.Seed(42)

	slice = make([]int, n)

	for i := 1; i < n; i++ {
		slice[i] = rand.Intn(256)
	}

	return
}

func branchPredictionIfElse(slice []int) (sum int) {
	var cnt int = len(slice)
	sum = 0

	for i := 0; i < cnt; i++ {
		if slice[i] >= 128 {
			sum += slice[i]
		}
	}

	return
}

func branchPredictionBitwise(slice []int) (sum int) {
	var cnt int = len(slice)
	sum = 0

	for i := 0; i < cnt; i++ {
		var t int = (slice[i] - 128) >> 31
		sum += (^t & slice[i])
	}

	return
}

func branchPredictionCompare() {
	const sliceLen int = 1 << 22
	var slice []int = initSlice(sliceLen)

	fmt.Printf("n: %v, if-else sum: %v\n", sliceLen, branchPredictionIfElse(slice))
	fmt.Printf("n: %v, bitwise sum: %v\n", sliceLen, branchPredictionBitwise(slice))
}
