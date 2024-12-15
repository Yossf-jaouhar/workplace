package functions

import "fmt"

func Average(nms []float64) float64 {
	if len(nms) == 0 {
		fmt.Println("No numbers")
		return 0
	}

	sum := 0.0
	for _, num := range nms {
		sum += num
	}

	return sum / float64(len(nms))
}
