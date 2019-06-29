package advanced

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// select blocks until one of its cases can run

// chooses one at random if multiple cases are ready

// select case can be:
// - assign (case var := <-channelName)
// - receiver (case <-channelName)
// - sender (case channelName <- "data")

// fallthrough is not available for select case

// you can dynamically add select cases with reflect.SelectCase

func textVomiter(
	textChannel chan []byte,
	quitChannel chan bool,
	waitGroup *sync.WaitGroup,
) {
	var counter int = 1
forLable:
	for {
		// first and second cases will be executed randomly
		select {
		case textChannel <- func() []byte {
			return []byte("testFromCase1_" + strconv.Itoa(counter))
		}():
			counter++
		case textChannel <- func() []byte { return []byte("testFromCase2") }():
			_ = true
		case <-quitChannel:
			fmt.Println("So Long, Suckers!")

			break forLable
		default:
			fmt.Println("Non-blocking operation...")
			time.Sleep(50 * time.Millisecond)
		}
	}

	waitGroup.Done()
}

// send-only channel typehint: chan<- Type
// receive-only channel typehint: <-chan Type
func channelSelect() {
	textChannel := make(chan []byte, 4)
	quitChannel := make(chan bool)

	// Used to ensure main goroutine waits
	// until all other is properly finish their work.
	var waitGroup *sync.WaitGroup = new(sync.WaitGroup)
	waitGroup.Add(2)

	// starting a sender
	go textVomiter(textChannel, quitChannel, waitGroup)

	// starting a receiver
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%s\n", <-textChannel)
		}

		quitChannel <- true

		waitGroup.Done()
	}()

	// main goroutine waits when both sender and receiver completes their jobs.
	waitGroup.Wait()
}

// Non-blocking operation...
// testFromCase2
// testFromCase2
// testFromCase1_1
// testFromCase1_2
// testFromCase2
// Non-blocking operation...
// testFromCase1_3
// testFromCase2
// testFromCase1_4
// testFromCase2
// testFromCase1_5
// So Long, Suckers!
