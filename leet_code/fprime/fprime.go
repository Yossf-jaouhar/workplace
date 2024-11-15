package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	if len(os.Args[1]) == 0 {
		return
	}

	for _, char := range os.Args[1]{
		if len(os.Args[1]) == 1 && char < '3' || char > '9'  {
			return
		}  else if  len(os.Args[1]) > 1 && char < '0' || char > '9' {
			return
		}
	}

	slice := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}
	nm, _ := strconv.Atoi(os.Args[1])

	result := ""

	for i := 0; i < len(slice); i++ {
		result += shhh(nm , slice[i])
	}
	if len(result) == 0 {
		fmt.Println(nm)
	}	
	if len(result) > 0 {
		result = result[:len(result)-1]
		fmt.Println(result)
	}
}


func shhh(nm int, pri int) string {
	result := ""
	for nm % pri == 0 {
		result += ch(pri)+"*"
		nm = nm / pri
	}
	return result
}

func ch(pri int)string{
	res := ""
	for pri != 0{
		r := pri % 10
		res = string(r+'0') + res
		pri = pri / 10
	}
	return res
}