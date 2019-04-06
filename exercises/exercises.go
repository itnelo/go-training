package exercises

import (
	"fmt"
	"unsafe"
)

func Run() {
	var testVal float64 = 1 << 20

	fmt.Printf("testVal: %v bytes, %v\n", unsafe.Sizeof(testVal), testVal)

	fmt.Printf("\nSquare root base prediction... %v\n", squareRootBaseEstimating(testVal))
	fmt.Printf("\nSquare root base prediction by formula... %v\n", squareRootBaseEstimatingFormula(testVal))
}
