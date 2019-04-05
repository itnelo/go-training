package advanced

import (
	"fmt"
	"math/rand"
)

const (
	CURRENT_PLATFORM_MAX_INT int = int(^uint(0) >> 1)
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
		var v int = slice[i]
		var t int = (v - 128) >> 31
		sum += (^t & v)
	}

	return
}

func branchPredictionCompare() {
	const sliceLen int = 1 << 22
	var slice []int = initSlice(sliceLen)

	fmt.Printf("Max int for current platform: %v\n", CURRENT_PLATFORM_MAX_INT)

	fmt.Printf("n: %v, if-else sum: %v\n", sliceLen, branchPredictionIfElse(slice))
	fmt.Printf("n: %v, bitwise sum: %v\n", sliceLen, branchPredictionBitwise(slice))
}
