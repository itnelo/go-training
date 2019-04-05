package basics

import (
	"fmt"
	"math/cmplx"
)

var x, y, z bool
var a, b int = 1, 2

var (
	isA    bool       = false
	maxInt int64      = 1<<63 - 1
	minInt int64      = -1 << 63
	cpx    complex128 = cmplx.Sqrt(-5 + 12i)
)

func typeConversions() {
	var i int = 42
	var fl float64 = float64(i)
	fmt.Println(fl, "is converted float64")

	var ui uint = uint(fl)
	fmt.Println(ui, "is converted uint")
}

func typeInference() {
	//v := 42
	//v := 42.1
	//v := "str1"
	v := 1.7 + 0.2i
	fmt.Printf("v is type %T\n", v)
}

func Variables() {
	var i int

	fmt.Println(x, y, z, i)

	fmt.Println(a, b)

	var foo, bar = true, "string"
	fmt.Println(foo, bar)

	k := 5
	fmt.Println(k)

	fmt.Println(isA, maxInt)

	fmt.Println(cpx)

	var _i int
	var _f float64
	var _b bool
	var _s string
	fmt.Printf("%v %v %v %q\n", _i, _f, _b, _s)

	typeConversions()

	typeInference()
}
