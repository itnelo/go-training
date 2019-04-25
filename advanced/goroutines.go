package advanced

import (
	"fmt"
	"runtime"
	_ "sync"
	"time"
)

// A goroutine is a lightweight thread managed by the Go runtime,
// because it uses little memory and resources and their initial stack size is small
// A real thread can hold many goroutines, but they can be within
// different threads if GOMAXPROCS > 1

// Starting from small amount and adjusting stack size dynamically.
// Shared address space, need to sync via channels.

// Goroutines runs only if main goroutine is running.

// Its all about concurrency, not parallelism.
// Concurrency vs Parallelism: https://golangbot.com/concurrency

// GOMAXPROCS is used for parallel execution (1 means it will be concurrent execution only).

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Println(s)
	}
}

func goroutines() {
	fmt.Println(runtime.GOMAXPROCS(0))

	go say("world")

	say("hello")
}
