package exercises

import (
	"fmt"
)

func Run() {
	var testVal float64 = float64(1 << 101)
	//fmt.Printf("\nSquare root base prediction... %v\n", squareRootBaseEstimating(testVal))
	fmt.Printf("\nSquare root base prediction by formula... %v\n", squareRootBaseEstimatingFormula(testVal))
}
