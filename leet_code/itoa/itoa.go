package main

import (
	"fmt"
	
)

func main() {
    fmt.Println(Itoa(12345))
    fmt.Println(Itoa(0))
    fmt.Println(Itoa(-1234))
    fmt.Println(Itoa(987654321))
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	nega := false
	if n < 0 {
		nega = true
		n = -n
	}
	res := ""
	for n != 0 {
		r := n % 10
		
		res += string(r + '0')
		n = n / 10
	}



	res1 := ""
	for i := len(res)-1; i >= 0 ; i--{
	
		res1 += string(int(res[i]))

	}

	if nega {
		res1 = string('-') + res1

	}

	return res1
}