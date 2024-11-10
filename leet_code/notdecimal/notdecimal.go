package main

import (
	"fmt"
)

func main() {
	fmt.Println(NotDecimal("0.0000000000000000000000001"))
	fmt.Println(NotDecimal("174.2"))
	fmt.Println(NotDecimal("0.0000000000000000001255"))
	fmt.Println(NotDecimal("1.205-25856"))
	fmt.Println(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Println(NotDecimal("-195.25856"))
	fmt.Println(NotDecimal("1952111"))
}

func NotDecimal(dec string) string {
	if len(dec) == 0 {
		return "\n"
	}
	s := 0
	if dec[0] == '-' {
		s =1
	}

	res := ""

	for i:=s ; i < len(dec) ; i++ {

		if dec[i] == '.' {
			continue
		} else if dec[i] < '0' || dec[i] > '9' {
			return dec 
		}
		
	

		res += string(dec[i]) 
	}



	c := 0
	for _,char := range res {
		if char != '0' {
			break
		}
		c++
	}


	res1 := ""
	for  i := c ; i < len(res) ; i++ {
		res1 += string(res[i])
	}

	return res1
}
