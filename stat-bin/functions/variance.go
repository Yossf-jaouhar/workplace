package functions

import "fmt"

func Variance(nms []float64, mean float64) float64 {
	if len(nms) == 0 {
		fmt.Println("no numbers")
		return 0
	}

	sum := 0.0
	for _, num := range nms {
		diff := num - mean 
		sum += diff * diff 
	}

	return sum / float64(len(nms))
}
