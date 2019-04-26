package basics

import (
	"fmt"
	"math"
	"runtime"
	"time"

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

func today_(t time.Weekday, offset int) time.Weekday {
	return t + time.Weekday(offset)
}

func dynamicSwitchExpression() bool {
	return true
}

// Note: "fallthrough" is not permitted in a type switch
// case [statement1], [statement2]: - is also works (stops at first positive check)
func switching() {
	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}

	fmt.Print("When's Saturday? ")

	var today time.Weekday = time.Now().Weekday()

	switch time.Saturday {
	case today_(today, 0):
		fmt.Println("Today")
	case today_(today, 1):
		fmt.Println("Tomorrow")
	case today_(today, 2):
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}

	t := time.Now()

	switch dynamicSwitchExpression() {
	case t.Hour() > 12:
		fmt.Println("t.Hour() > 12")
	default:
		fmt.Println("t.Hour() <= 12")
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

func typeSwitch(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("an integer")
	case string:
		fmt.Println("a string")
	// v is the same type as i
	default:
		fmt.Printf("(%v, %T)\n", v, v)
	}
}

func Controlflow() {
	//forloop()
	forInvariantOnly()

	fmt.Printf("limited pow: %v\n", ifScopeStatement(2, 10, 1<<10-1))

	var reversedStr string = stringutil.ReverseRange("!oG ,olleH")
	fmt.Println("ReverseRange: " + reversedStr)

	reversedStr2 := stringutil.ReverseConvert("2 !oG ,olleH")
	fmt.Println("ReverseConvert: " + reversedStr2)

	fmt.Println("ReverseDeferred: " + stringutil.ReverseDeferred("2 !oG ,olleH"))

	switching()
	typeSwitch(5.55)
}
