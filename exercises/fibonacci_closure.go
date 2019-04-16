package exercises

import (
	"fmt"
)

func fibonacci(debug bool) func() int {
	var prev, current int = 1, 0

	// defer doesn't rewrite a retval because it is unnamed here
	return func() int {
		// added on stack, will be executed before caller can see a retval
		defer func() {
			prev, current = current, prev+current

			if debug {
				fmt.Println("new pair generated:", prev, current)
			}
		}()

		if debug {
			fmt.Println("returning...", current)
		}

		return current
	}
}

func fibonacciClosure() {
	f := fibonacci(false)

	for i := 0; i < 10; i++ {
		fmt.Print(f(), " ")
	}

	fmt.Print("\n\n")

	// debug true
	f2 := fibonacci(true)

	for i := 0; i < 10; i++ {
		fmt.Print("received by caller: ", f2(), " \n\n")
	}
}
