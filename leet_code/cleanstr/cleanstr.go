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

	mmap := make(map[rune]int)
	mmap1 :=  make(map[rune]int)

	for _, char := range s1 {
		mmap[char] = 1
	}
	


	res3 := ""
	for _ , st := range s2 {
		mmap1[st] = 1
		if mmap[st] == 1 {
			res3 += string(st)
			 mmap[st] = 0
		}
	}
	
	fmt.Println(res3)
}
