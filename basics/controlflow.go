package basics

import (
	"fmt"
	"math"

	"github.com/itnelo/stringutil"
)

func forloop() {
	var sum int8 = 5

	for i := int8(1); i < 10; i++ {
		sum := sum + i

		fmt.Println(sum)
	}
}

func forInvariantOnly() {
	var i int = 1

	for i <= 10 {
		if i%2 == 0 {
			fmt.Println(i)
		}

		i++
	}
}

func infinityLoop() {
	for {

	}
}

func ifScopeStatement(x, n, limit float64) float64 {
	if v := math.Pow(x, n); v < limit {
		return v
	} else {
		fmt.Printf("%v > %v, so we use the limit as a retval\n", v, limit)
	}

	// v is undefined here
	//fmt.Printf("pow from local if scope: %v\n", v)

	return limit
}

func Controlflow() {
	//forloop()
	forInvariantOnly()

	fmt.Printf("limited pow: %v\n", ifScopeStatement(2, 10, 1<<10-1))

	var reversedStr string = stringutil.ReverseRange("!oG ,olleH")
	fmt.Println("ReverseRange: " + reversedStr)

	reversedStr2 := stringutil.ReverseConvert("2 !oG ,olleH")
	fmt.Println("ReverseConvert: " + reversedStr2)
}
