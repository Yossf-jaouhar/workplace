package functions

import "fmt"

func Median(nms []int) int {
	if len(nms) == 0 {
		fmt.Println("No numbers to calculate median.")
		return 0
	}

	res := Quicksort(nms)

	mid := len(res) / 2
	if len(res)%2 == 0 {
		return (res[mid-1] + res[mid]) / 2
	}
	return res[mid]
}

func Quicksort(nms []int) []int {
	if len(nms) <= 1 {
		return nms
	}

	key := nms[0]

	left := []int{}
	right := []int{}

	for i := 1; i < len(nms); i++ {
		if key > nms[i] {
			left = append(left, nms[i])
		} else {
			right = append(right, nms[i])
		}
	}
	return append(append(Quicksort(left), key), Quicksort(right)...)
}
