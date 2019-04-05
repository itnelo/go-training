package basics

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 3 / 10
	y = sum - x
	return
}

func Functions() {
	fmt.Println(add2(add(42, 13), 5))

	a := "Hello"
	b := "World"
	fmt.Println(a, b)

	a, b = swap(a, b)
	fmt.Println(a, b, "swapped")

	fmt.Println(split(100))
}
