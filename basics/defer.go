package basics

import (
	"fmt"
)

// 1. For clean-up tasks (mutex unlocks, closing connections/resources, etc.)
// 2. Recovering from panic
// 3. Managing error codes, etc.

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

func oofPanic(i int) {
	if i > 3 {
		fmt.Println("Panicking!")

		panic(fmt.Sprintf("Panic at %v iteration.", i))
	}

	defer fmt.Printf("Defer in oofPanic, %v iteration\n", i)
	fmt.Printf("Printing in oofPanic, %v iteration\n", i)

	oofPanic(i + 1)
}

func deferPanicRecovering() (error string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic captured:", r)

			if rStr, isString := r.(string); isString {
				error = "An error has been occurred during function call: " + rStr
			} else {
				// Process unrecoverable situation.
				fmt.Println("Unrecoverable situation.")
			}

			return
		}
	}()

	fmt.Println("Calling function that can panic.")

	oofPanic(0)

	fmt.Println("Continue execution flow normally.")

	return
}

func Deferred() {
	printDeferred()
	deferredCounting()

	//deferExecutedAfterPanic()
	var error string = deferPanicRecovering()

	fmt.Printf("Error content from deferPanicRecovering() func: %v\n", error)

	//Calling function that can panic.
	//Printing in oofPanic, 0 iteration
	//Printing in oofPanic, 1 iteration
	//Printing in oofPanic, 2 iteration
	//Printing in oofPanic, 3 iteration
	//Panicking!
	//Defer in oofPanic, 3 iteration
	//Defer in oofPanic, 2 iteration
	//Defer in oofPanic, 1 iteration
	//Defer in oofPanic, 0 iteration
	//Panic captured: Panic at 4 iteration.
	//Error content from deferPanicRecovering() func: An error has been occurred during function call: Panic at 4 iteration.
}
