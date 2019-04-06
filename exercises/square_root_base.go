package exercises

import (
	"fmt"
	"math"
	"unsafe"
)

const (
	FACTOR   float64 = 2
	ACCURACY float64 = 1 / 1000000
)

func squareRootBaseEstimating(x float64) float64 {
	var estimate float64

	for {
		if xSqrt := math.Pow(estimate+1, FACTOR); xSqrt <= x {
			estimate += 1
		} else {
			break
		}
	}

	return estimate
}

func squareRootBaseEstimatingFormula(x float64) float64 {
	var estimate float64 = 1.0

	for math.Abs(x-estimate*estimate) > ACCURACY {
		estimate -= (estimate*estimate - x) / (estimate * 2)
	}

	return estimate
}

func squareRootBaseEstimatingBuiltin(x float64) float64 {
	return math.Sqrt(x)
}

func evalSquareRootBaseFunctions() {
	var testVal float64 = 1 << 20

	fmt.Printf("testVal: %v bytes, %v\n", unsafe.Sizeof(testVal), testVal)

	fmt.Printf("\nSquare root base prediction... %v\n", squareRootBaseEstimating(testVal))
	fmt.Printf("\nSquare root base prediction by formula... %v\n", squareRootBaseEstimatingFormula(testVal))
}
