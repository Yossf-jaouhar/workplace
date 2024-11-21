package main

import "fmt"

func main() {
	nm1 := []int{5, 2}
	nm2 := []int{3, 4, 2, 5, 1}
	n := findMedianSortedArrays(nm1, nm2)
	fmt.Println(n)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	var res []int
	i,j := 0 , 0

	for i < len(nums1) && j < len(nums2) {

		if nums1[i] < nums2[j]{
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}

	for i < len(nums1) {
		res = append(res, nums1[i])
		i++
	}


	for j < len(nums2) {
		res = append(res, nums2[j])
		j++
	}

	
	fmt.Println(res)
	var v float64
	if len(res) % 2 == 0 {
		 b :=  float64(res[len(res) / 2]) + float64(res[(len(res) / 2)-1])
		 v = b / 2.0
	} else {
	 v = float64(res[len(res) / 2]) 
	}
	return v
}