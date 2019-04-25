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

type Context struct {
	data       []int
	lowerBound int
	upperBound int
}

func calculateSum(context *Context, c chan int) {
	fmt.Printf(
		"Starting calculation for data[%d:%d]\n",
		context.lowerBound,
		context.upperBound,
	)

	var sum int = 0
	var partialData = context.data[context.lowerBound:context.upperBound]

	for _, value := range partialData {
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

	// unbuffered channel
	// var results = make(chan int)

	// channel with buffer capacity
	var resultsChannel = make(chan int, cpuCount)

	for i := 0; i < cpuCount; i++ {
		var lowerBound, upperBound int = i * batchSize, (i + 1) * batchSize

		go calculateSum(&Context{data, lowerBound, upperBound}, resultsChannel)
	}

	var sum int

	for i := 0; i < cpuCount; i++ {
		var partialSum = <-resultsChannel

		fmt.Println("Partial sum received:", partialSum)

		sum += partialSum
	}

	// Using range to continuously fetch response from the channel.
	// Note: we need to close() somewhere before or within worker context
	// to properly stop this loop.
	//
	// for partialSum := range resultsChannel {
	// 	//...
	// }

	close(resultsChannel)

	fmt.Println("Sum:", sum)
}
