package main

import (
	"github.com/itnelo/go-training/advanced"
	"github.com/itnelo/go-training/basics"
	"github.com/itnelo/go-training/exercises"
	_ "github.com/symfony-doge/event/example"
	_ "github.com/symfony-doge/splitex/example"
)

func main() {
	basics.Run()
	advanced.Run()
	exercises.Run()

	// github.com/symfony-doge/event/example
	// example.OneSubscriberManyPublishers()

	// github.com/symfony-doge/splitex/example
	// example.ConcurrentSliceSum()
}
