package functions

import "math"

func StandardDeviation(variance float64) float64 {
	deviation := math.Sqrt(variance)
	return deviation
}
