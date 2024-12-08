package functions

import "fmt"

func Average(nms []int) int {
	if len(nms) == 0 {
		fmt.Println("No numbers")
		return 0
	}

	sum := 0
	for _, num := range nms {
		sum += num
	}

	return sum / len(nms)
}
