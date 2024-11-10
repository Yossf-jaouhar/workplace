package main

import (
	"fmt"
	
)

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}



func Atoi(s string) int {

	start := 0
	nega := false

	if s[0] == '-' {
		start =1
		nega = true
	}
	if s[0] == '+' {
		start = 1
	}


	res := 0
	for i := start ; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			res = res * 10 + int(s[i]- '0')
		
		} else {
			return 0
		}
	}

	if nega {
		res = - res
	}
	return res
}