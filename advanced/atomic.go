package advanced

import (
	"fmt"
	"sync/atomic"
	"time"
)

func atomicIncrement() {
	var shared uint64

	for i := 0; i < 50; i++ {
		go func() {
			atomic.AddUint64(&shared, 1)
		}()
	}

	time.Sleep(100 * time.Microsecond)

	sharedResult := atomic.LoadUint64(&shared)

	fmt.Println("sharedResult:", sharedResult)
}
