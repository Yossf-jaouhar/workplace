package main

import "fmt"

func main() {
	nm := mySqrt(8)
	fmt.Println(nm)
}

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	r := x

	for r*r> x {
		r = (r + x/r) / 2
	}

	return r
}
