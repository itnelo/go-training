package main

import (
	"fmt"
)

var x, y, z bool
var a, b int = 1, 2

var (
	isA    bool  = false
	maxInt int64 = 1<<63 - 1
)

func variables() {
	var i int

	fmt.Println(x, y, z, i)

	fmt.Println(a, b)

	var foo, bar = true, "string"
	fmt.Println(foo, bar)

	k := 5
	fmt.Println(k)

	fmt.Println(isA, maxInt)
}
