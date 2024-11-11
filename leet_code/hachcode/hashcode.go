package main

import (
	"fmt"
)

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}



func HashCode(dec string) string {

	
	ress := ""
	for _, char := range dec {
		res := (int(char) + len(dec)) % 127
		if res < 32 && res > 126 {
			res = res +33
		}
		ress += string(res)

	}
	return ress
}