package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]
	res := ""

	c := 0
	
	for _, st := range s1{

		for i := c; i < len(s2) ; i++{
			if byte(st) == s2[i] {
				res += string(st)
				c = i
				break
			}
		}
	}
	if len(s1) == len(res) {
		fmt.Println(res)
	} else {
		return
	}
}
