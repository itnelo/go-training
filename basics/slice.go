package basics

import (
	"fmt"
)

func slices() {
	s := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%v, len(s) == %v, cap(s) == %v\n", s, len(s), cap(s))

	a := make([]int, 5, 5)

	fmt.Printf("%v, len(a) == %v, cap(a) == %v, addr == %p\n", a, len(a), cap(a), &a)

	a = append(a, 1)
	fmt.Printf("%v, len(a) == %v, cap(a) == %v, addr == %p\n", a, len(a), cap(a), &a)
	a = append(a, 1)
	fmt.Printf("%v, len(a) == %v, cap(a) == %v, addr == %p\n", a, len(a), cap(a), &a)

	// todo
}
