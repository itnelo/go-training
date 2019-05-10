package advanced

import (
	"fmt"
	"time"
)

// time.Timer is a single-time trigger
// time.Ticker is a repetitive trigger

// use-cases comparison against time.Sleep:
// - can be stopped
// - can be used as a "bursty rate limiter" if have a buffer
//     (see https://gobyexample.com/rate-limiting)

// Some kind of data for state monitoring.
type Resource struct {
	name string
	data int
}

// State definition.
type State struct {
	code        int
	description string
}

// Returns a channel for pushing updates which will be analysed by a monitor.
// Note: we using here a send-only typehint for a channel-retval.
// If we use named retvals we cannot receive from this channel
// within local scope (see below).
func stateMonitorInit(updateInterval time.Duration) chan<- *Resource {
	// It is a common channel without any receive/send limitations.
	var analyseChannel = make(chan *Resource)

	var ticker = time.NewTicker(updateInterval)
	var timer = time.NewTimer(updateInterval)
	var stateMap = make(map[string]State)
	var running = true

	go func() {
		for running {
			select {
			case time := <-timer.C:
				fmt.Println(time, "Timer alarm")
			case time := <-ticker.C:
				fmt.Println(time, "Resource state:", stateMap)
			// We can read from this channel, no limitations.
			case resource, isChannelOpen := <-analyseChannel:
				if !isChannelOpen {
					fmt.Println("Channel is closed, stopping select loop...")

					running = false

					break
				}

				fmt.Println("Resource received for analysis:", resource)

				stateMap[resource.name] = analyseResource(resource)
			}
		}

		fmt.Println("Exiting analyser goroutine...")
	}()

	// Now we allow consumers to only send to this channel because retval
	// typehinted as <-chan.
	return analyseChannel
}

func analyseResource(resource *Resource) State {
	var code, description = 500, "FAIL"

	if resource.data > 0 {
		code, description = 200, "OK"
	}

	return State{
		code:        code,
		description: description,
	}
}

func tickerAndTimer() {
	var resource = &Resource{name: "resource1", data: 0}

	analyseChannel := stateMonitorInit(3 * time.Second)

	for range []int{1, 2, 3} {
		fmt.Println("Sending resource to the analysis channel...")

		analyseChannel <- resource
	}

	time.Sleep(10 * time.Second)

	resource.data = 1
	analyseChannel <- resource

	time.Sleep(10 * time.Second)

	close(analyseChannel)

	time.Sleep(2 * time.Second)
}

// Sending resource to the analysis channel...
// Resource received for analysis: &{resource1 0}
// Sending resource to the analysis channel...
// Sending resource to the analysis channel...
// Resource received for analysis: &{resource1 0}
// Resource received for analysis: &{resource1 0}
// 2019-05-10 12:59:53.288531994 +0300 MSK m=+3.001080175 Resource state: map[resource1:{500 FAIL}]
// 2019-05-10 12:59:53.288541102 +0300 MSK m=+3.001089250 Timer alarm
// 2019-05-10 12:59:56.288928082 +0300 MSK m=+6.001476255 Resource state: map[resource1:{500 FAIL}]
// 2019-05-10 12:59:59.288601144 +0300 MSK m=+9.001149311 Resource state: map[resource1:{500 FAIL}]
// Resource received for analysis: &{resource1 1}
// 2019-05-10 13:00:02.288739267 +0300 MSK m=+12.001287496 Resource state: map[resource1:{200 OK}]
// 2019-05-10 13:00:05.288956285 +0300 MSK m=+15.001504515 Resource state: map[resource1:{200 OK}]
// 2019-05-10 13:00:08.288581512 +0300 MSK m=+18.001129734 Resource state: map[resource1:{200 OK}]
// Channel is closed, stopping select loop...
// Exiting analyser goroutine...
