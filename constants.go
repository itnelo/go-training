package main

import (
	"fmt"
)

const (
	TEST_CNST      int    = 7
	TEST_CNST2     string = "test1"
	CHAR_CNST      byte   = 'a'
	MAX_ASCII_CHAR byte   = 1<<8 - 1
	RUNE_CNST      rune   = 0x2318 // '\u2318'

	RUNE_STRING string = "пробираюсь через дремучий лес..."

	BIG   = 1 << 100
	SMALL = BIG >> 99
)

func needInt(x int) (y int) {
	y = x*10 + 1

	return
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func numericConstantsHighPrecision() {
	fmt.Println(needInt(SMALL))
	//fmt.Println(needInt(BIG))
	fmt.Println(needFloat(SMALL))
	fmt.Println(needFloat(BIG))
}

func constants() {
	const INLINE_CONST = 1.5 + 0.9i

	fmt.Println(TEST_CNST, TEST_CNST2, INLINE_CONST)

	fmt.Printf("max ascii char: %c\n", MAX_ASCII_CHAR)
	fmt.Printf("char: %c, rune: %c\n", CHAR_CNST, RUNE_CNST)

	for index, runeValue := range RUNE_STRING {
		fmt.Printf("%#U starts at byte position %v\n", runeValue, index)
	}

	numericConstantsHighPrecision()
}
