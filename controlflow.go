package main

import (
	"fmt"
)

func forloop() {
	var sum int8 = 5

	for i := int8(1); i < 10; i++ {
		sum := sum + i

		fmt.Println(sum)
	}
}

func controlflow() {
	forloop()
}
