package main

import (
	"fmt"
	"os"
)


func main(){
	if len(os.Args) > 2  || len(os.Args) == 1	{
		return 
	}

	res := ""

	for _ , char := range os.Args[1] {
		
		if char >= 'a' && char <= 'z' {
		
			if char + 13 <= 'z' {
				res += string(char + 13)
			} else if  char + 13 > 'z' {
				res += string(char - 13)
			}
		} else if char >= 'A' && char <= 'Z' {
			if char + 13 <= 'Z' {
				res += string(char + 13)
			} else if  char + 13 > 'Z' {
				res += string(char - 13)
			}
		} else {
			res += string(char)
		}
	}

	fmt.Println(res)

}