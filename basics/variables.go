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

func strings() {
	// string is a character (byte aka uint8) sequence
	var str string = "пYривет"
	// [0:2] == "п" (read 2 byte rune / int32)
	// [2:3] == "Y" (read 1 byte ASCII)
	// [3:5] == "р" (read 2 byte rune / int32)
	// etc.

	// index access should be by slice in case of UTF-8
	fmt.Printf("%s\n", str[3:5])

	// best practices:
	// range yields (index, tune) from string
	// []rune("привет")
	// []int32("привет")
	// []byte("test") - ASCII
	// []uint8("test") - ASCII
	// string(byte[]("..."))
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

	strings()
}
