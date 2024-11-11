package main

import "fmt"



func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintMemory(arr [10]byte) {

	bs := "0123456789ABCDEF"


	res := ""

	c := 1
	for _,char :=range arr {
	
		if  c == 4 {

			
			res += string(bs[int(char)/16]) +  string(bs[int(char) % 16])
			
			fmt.Println(res)
			res = ""
			c = 1
		} else {

			if char == 0 {
				res += "00"
			}
			
			res += string(bs[char/16]) +  string(bs[char % 16]) + " "
			c++


		}

		
	}

	fmt.Println()

	// for _, char := range arr {

	// }
}