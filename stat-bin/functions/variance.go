package functions

import "fmt"

func Variance(nms []int, mean int) int {
	if len(nms) == 0 {
		fmt.Println("No numbers to calculate variance.")
		return 0
	}

	sum := 0
	for _, num := range nms {
		sum += (num - mean) * (num - mean)
	}	

	return sum / len(nms)
}
