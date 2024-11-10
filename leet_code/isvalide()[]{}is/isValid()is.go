package main

import "fmt"

func main() {
	s := "[]"
	fmt.Println(isValid(s)) 
}	



func isValid(s string) bool {
    var gg []rune


    for _, char := range s{

		switch char {
		case  '(','{','[' :	
			gg = append(gg, char)
		case ')':
			if len(gg) == 0 || gg[len(gg)-1] != '(' {
				return false
			}
			gg = gg[:len(gg)-1]
		case ']':
			if len(gg) == 0 || gg[len(gg)-1] != '[' {
				return false
			}
			gg = gg[:len(gg)-1]
		case '}':
			if len(gg) == 0 || gg[len(gg)-1] != '{' {
				return false
			}
			gg = gg[:len(gg)-1]	
		}
	}
	return len(gg) == 0
}
