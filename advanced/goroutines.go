package advanced

import (
	"fmt"
	_ "sync"
	"time"
)

// A goroutine is a lightweight thread managed by the Go runtime
// starting from small amount and adjusting stack size dynamically
// same address space

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Println(s)
	}
}

func goroutines() {
	go say("world")

	say("hello")
}
