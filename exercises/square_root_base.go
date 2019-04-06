package exercises

import (
	"math"
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
