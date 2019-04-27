package exercises

import (
	"fmt"
	"runtime"
	"time"
)

// ping pong concurrency pattern
// https://talks.golang.org/2013/advconc.slide#6

type Ball struct {
	hits int
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++

		fmt.Println(name, ball.hits)

		time.Sleep(100 * time.Millisecond)

		table <- ball
	}
}

func pingPong() {
	fmt.Println(runtime.GOMAXPROCS(0))

	table := make(chan *Ball)

	go player("pong", table)
	go player("ping", table)

	table <- new(Ball) // game on; toss the ball

	time.Sleep(1 * time.Second)

	<-table // game over; grab the ball
}
