package main

import "fmt"

func selectionSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {

		minIdx := i
		for j := i + 1; j < n; j++ {
			fmt.Println(j)
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
}

func main() {
	numbers := []int{64, 25, 12, 22, 11}
	

	selectionSort(numbers)
	//fmt.Println("after:", numbers)
}
