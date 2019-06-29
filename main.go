package main

import (
	"github.com/itnelo/go-training/advanced"
	"github.com/itnelo/go-training/basics"
	"github.com/itnelo/go-training/exercises"
	"github.com/symfony-doge/event/example"
)

func main() {
	basics.Run()
	advanced.Run()
	exercises.Run()

	// github.com/symfony-doge/event/example
	example.OneSubscriberManyPublishers()
}
