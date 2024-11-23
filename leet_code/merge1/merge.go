package main

import "fmt"

func main() {
	// Example 1
	nums1 := []int{1, 2, 0, 0, 0, 0, 0}
	m := 2
	nums2 := []int{2, 6, 5, 6, 8}
	n := 5

	// Calling the merge function
	merge(nums1, m, nums2, n)

	// Output the merged nums1
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
	var res []int
	
	res =append(res, nums1...)
	res = append(res, nums2...)
	ress := sort(res)
	fmt.Println(ress)
}

func sort(n1 []int) []int {
	a := len(n1) / 2
	left := sort(n1[:a])
	right := sort(n1[a:])

	res := add(left, right)
	return res
}
func add(l , r []int) []int {
	var res []int
	i, j := 0,0

	for i < len(l) {
		if l[i] < r[j] {
			res =append(res, l[i])
		    i++
		} else {
			res =append(res, l[j])
		    j++
		}
		
	}
	if i < len(l) {
		res = append(res, l[i])
		i++
	}
	if j < len(r) {
		res = append(res, r[j])
		j++
	}
	return res
}