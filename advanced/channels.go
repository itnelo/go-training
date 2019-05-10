package advanced

import (
	"fmt"
	"runtime"
	"unsafe"
)

// builtin typed bridge between goroutines
// sync goroutines without any locks or condition variables
// By default <- blocks until sender becomes ready

// <- channel operator for receiving data
// deadlock:
// - nil channel
// - no reader/writer

// conventions:
// - communicate less, do more
// - only sender closes the channel
// - closing a channel is optionally, only if an explicit "no more values" flag is needed
// - use send-only (chan<- Type) and receive-only (<-chan Type) typehints for additional type-safety

// We can organize receive op for a channel within local scope
// but restrict consumer that gets this channel by func retval
// to only sending op (see ticker_and_timer.go example)

// messaging between goroutines is pricey, especially if they executed by
// different OS threads

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

	// Blocks execution of goroutine if channel is full
	// (until some receiver will read it)
	c <- sum
}

func channels() {
	var data []int = make([]int, DataAmount)

	for i := 0; i < DataAmount; i++ {
		data[i] = 1
	}

	// runtime.GOMAXPROCS is responsible for parallelism
	// 1 means it is only concurrent execution
	fmt.Println("runtime.GOMAXPROCS ==", runtime.GOMAXPROCS(0))
	//runtime.GOMAXPROCS(1)

	var cpuCount int = runtime.NumCPU()
	var dataBytes int = int(unsafe.Sizeof(cpuCount)) * DataAmount

	fmt.Printf("Processing %d bytes of data using %d CPU\n", dataBytes, cpuCount)

	var batchSize int = DataAmount / cpuCount

	// unbuffered channel
	// sends only if at least one receiver is awaits
	// var results = make(chan int)

	// Channel with buffer capacity
	// Means we can write N integers in this channel before it blocks
	// If the writes on channel are more than its capacity,
	// then the writes are not processed till its concurrent reading is done
	// from one of the goroutines, and once that is done,
	// it will write new values to the channel.
	var resultsChannel chan int = make(chan int, cpuCount)
	defer close(resultsChannel)

	for i := 0; i < cpuCount; i++ {
		var lowerBound, upperBound int = i * batchSize, (i + 1) * batchSize

		go calculateSum(&Context{data, lowerBound, upperBound}, resultsChannel)
	}

	var sum int = 0

	for i := 0; i < cpuCount; i++ {
		// Blocks execution of main goroutine if channel is empty
		// (value, ok), where ok indicates that channel is not closed yet
		partialSum, isChannelStillOpen := <-resultsChannel

		if !isChannelStillOpen {
			break
		}

		fmt.Println("Partial sum received:", partialSum)

		sum += partialSum
	}

	// Using range to continuously fetch response from the channel.
	// Note: we need to close() within worker (sender) context
	// to properly stop this loop.
	//
	// for partialSum := range resultsChannel {
	// 	//...
	// }

	fmt.Println("Sum:", sum)
}

// runtime.GOMAXPROCS == 8
// Processing 262144 bytes of data using 8 CPU
// Starting calculation for data[28672:32768]
// Partial sum received: 4096
// Starting calculation for data[0:4096]
// Partial sum received: 4096
// Starting calculation for data[24576:28672]
// Partial sum received: 4096
// Starting calculation for data[4096:8192]
// Starting calculation for data[8192:12288]
// Starting calculation for data[12288:16384]
// Starting calculation for data[20480:24576]
// Partial sum received: 4096
// Partial sum received: 4096
// Partial sum received: 4096
// Partial sum received: 4096
// Starting calculation for data[16384:20480]
// Partial sum received: 4096
// Sum: 32768
