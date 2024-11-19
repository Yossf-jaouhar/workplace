package main

import "fmt"

func main() {
	nm1 := []int{1, 6, 5, 2}
	nm2 := []int{3, 4, 2, 5}
	n := findMedianSortedArrays(nm1, nm2)
	fmt.Println(n)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var res []int
	for i :=0; i < len(nums1); i++ {
		for j := 0; j < len(nums1)-i-1 ; j++ {
			if nums1[j] > nums1[j+1] {
				nums1[j], nums1[j+1] =  nums1[j+1], nums1[j] 
			}
		}
	}
    res = append(res, nums1...)
    for i :=0; i < len(nums2); i++ {
		for j := 0; j < len(nums2)-i-1 ; j++ {
			if nums2[j] > nums2[j+1] {
				nums2[j], nums2[j+1] =  nums2[j+1], nums2[j] 
			}
		}
	}

    res = append(res, nums2...)

    for i :=0; i < len(res); i++ {
		for j := 0; j < len(res)-i-1 ; j++ {
			if res[j] > res[j+1] {
				res[j], res[j+1] =  res[j+1], res[j] 
			}
		}
	}


    var v float64

    if len(res) / 2 == 0 {
        v = float64(res[len(res)/2]) + float64(res[(len(res)/2)+1])
        v = v /2
    } else {
        v = float64(res[len(res)/2])
    }

    fmt.Println(res)
    return v
}
