package main

import (
	"fmt"
	"unsafe"
)

func lowlevel() {
	var x uint8 = 24

	//24 / 2 = 12 (0)
	//12 / 2 = 6  (0)
	//6 / 2 = 3   (0)
	//3 / 2 = 1   (1)
	//        |__ (1)

	// 011000 = 1 * 2^4 + 1 * 2^3 = 16 + 8 = 24
	fmt.Printf("x(24) in binary: %b", x)

	// obtain memory address of variable "x"
	x_memaddr := unsafe.Pointer(&x)

	fmt.Printf("\nmem addr for x: %#v", x_memaddr)

	// copying value by memory address
	var y *uint8

	fmt.Printf("\ny == nil: %v", y == nil)

	y = (*uint8)(x_memaddr)

	// 24 / 16 = 1  (8)
	//           |_ (1)
	// 0x18
	fmt.Printf("\ncopied y value in hex: %#x", *y)

	var xx uint16 = 1<<16 - 1
	var xx_memaddr (unsafe.Pointer) = unsafe.Pointer(&xx)
	var b *byte
	b = (*byte)(xx_memaddr)

	fmt.Printf("\nunsafe.Pointer carries no type information, so value has been truncated: %v", *b)
}
