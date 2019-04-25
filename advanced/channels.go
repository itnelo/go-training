package advanced

import (
	"fmt"
	"runtime"
	"unsafe"
)

// builtin typed bridge between goroutines
// sync goroutines without any locks or condition variables
// By default <- blocks until sender becomes ready

// <- channel operator for sync

// conventions: communicate less, do more
// messaging between goroutines is pricey

const (
	DataAmount int = 1 << 15
)

func calculate(data []int, c chan int) {
	fmt.Println("Starting calculation for data banch", len(data))

	var sum int = 0

	for _, value := range data {
		sum += value
	}

	c <- sum
}

func channels() {
	var data = make([]int, DataAmount)

	for i := 0; i < DataAmount; i++ {
		data[i] = 1
	}

	var cpuCount int = runtime.NumCPU()
	var dataBytes = int(unsafe.Sizeof(cpuCount)) * DataAmount

	fmt.Printf("Processing %d bytes of data using %d CPU\n", dataBytes, cpuCount)

	var batchSize int = DataAmount / cpuCount
	var partialSum = make(chan int)

	for i := 0; i < cpuCount; i++ {
		var lowBound, highBound int = i * batchSize, (i + 1) * batchSize
		go calculate(data[lowBound:highBound], partialSum)
	}

	var sum int

	for i := 0; i < cpuCount; i++ {
		var x int = <-partialSum

		fmt.Println("Partial data received:", x)

		sum += x
	}

	close(partialSum)

	fmt.Println("Sum:", sum)
}
