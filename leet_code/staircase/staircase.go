package main

import "fmt"

func main() {
	v := staircase(6)
	fmt.Println(v)
}

func staircase(n int) string {
	b := n-1
	v := n
	res := ""
	for i := 0; i < v; i++ {

		for j := 0; j < v; j++ {
			if b <= j {
				res += "#"
			} else {
				res += " "
			}
		}
		b--

		if i+1 == n {
			continue
		} else {
			res += "\n"
		}
	}
	return res
}
