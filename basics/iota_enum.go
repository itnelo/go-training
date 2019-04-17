package basics

import (
	"fmt"

	_ "unsafe" // testing go:linkname
)

// https://golang.org/ref/spec#Iota

type Level uint16
type Fruit uint8
type Bits uint8

const (
	Apple Fruit = iota // + 1 => 1, 2, 3
	Pineapple
	Banana
	Test
)

const (
	Zero  Level  = iota // 0
	One          = iota // 1
	Two          = 2    // 2
	Three        = iota // 3
	_                   // omitting 4
	Five  = iota        // 5
)

const (
	B0 Bits = 1 << iota
	B1
	B2
)

// pairs
const (
	a1, b1 = iota, iota * 2 // 0, 0
	a2, b2                  // 1, 2
	a3, b3                  // 2, 4
	a4, b4                  // 3, 6
)

var fruits = [...]string{
	"Apple",
	"Pineapple",
	"Banana",
}

func (f Fruit) String() string {
	// check bounds here
	if Apple > f || f > Banana {
		panic("Fruit doesn't exists.")
	}

	return fruits[f]
}

func iotaEnumerations() {
	fmt.Println(Zero, One, Two, Three, Five)

	fmt.Println(Apple, Pineapple, Banana, Test)
	fmt.Printf("typeof Apple: %T\n", Apple)

	var binaryMask Bits
	binaryMask |= B1
	binaryMask |= B2
	fmt.Printf("binaryMask (|= B1, |= B2): %b\n", binaryMask)
	binaryMask &^= B1 // clear bit
	fmt.Printf("binaryMask (&^= B1): %b\n", binaryMask)
	fmt.Println("binaryMask has B1?", binaryMask&B1 != 0)

	fmt.Println("Pair (a2, b2):", a3, b3)
}
