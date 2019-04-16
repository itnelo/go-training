package basics

import (
	"fmt"
)

// lambda - is an anonymous function
// closure - a lambda that can access variables from enclosuring context (i.e. visibility block)

func Incrementor() func(int, int) int {
	// sum is not gc-ed while it is in usage (bounded with closure)
	sum := 0

	// second value is unused (use case: interfaces and different implementations without last arg requirement)
	return func(x, _ int) int {
		sum += x

		return sum
	}
}

func closures() {
	// each closure has its own context instance
	closure := Incrementor()
	closure2 := Incrementor()

	closure2(100, 10)

	for i := 0; i < 10; i++ {
		fmt.Print(closure(i, 5), ":", closure2(i, 6), " ")
	}

	fmt.Println()
}
