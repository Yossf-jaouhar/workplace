package main

import (
	"fmt"
)

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}


func FromTo(from int, to int) string {

	if from < 0 || from > 99 {
		return "Invalid\n"
	}


	if from == to {
		return string(to)
	}


	res := ""
	if from < to {
		for i:= from; i <= to; i++ {
		    res +=  string((i/10) + '0') + string((i%10) + '0')
            if i < to {
				res += ", "
			}
		}
		
	}


	if to < from {
		for i:= to; i <= from; i++ {
			res += string((i/10)+ '0') + string((i%10) + '0')
			if i < from {
				res += ", "
			}
		}
		
	}

	return res + "\n"

}