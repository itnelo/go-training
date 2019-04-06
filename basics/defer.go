package basics

import (
	"fmt"
)

func printDeferred() {
	defer fmt.Println("world")

	fmt.Println("Hello")
}

func deferredCounting() {
	fmt.Println("Counting...")

	for i := 0; i < 10; i++ {
		defer fmt.Print(i, " ")
	}

	fmt.Println("\nDone.")
}

func deferExecutedAfterPanic() {
	defer fmt.Println("I'am normally executed.")

	panic("Need more toilet paper!")
}

func Deferred() {
	printDeferred()
	deferredCounting()

	deferExecutedAfterPanic()
}
