package main

import (
	"fmt"

	"github.com/itnelo/stringutil"
)

func forloop() {
	var sum int8 = 5

	for i := int8(1); i < 10; i++ {
		sum := sum + i

		fmt.Println(sum)
	}
}

func controlflow() {
	//forloop()

	var reversedStr string = stringutil.ReverseRange("!oG ,olleH")
	fmt.Println("ReverseRange: " + reversedStr)

	reversedStr2 := stringutil.ReverseConvert("2 !oG ,olleH")
	fmt.Println("ReverseConvert: " + reversedStr2)

}
