package main

import (
	"fmt"
)

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}


func ZipString(s string) string {

	cont := 1
	res := ""
	for i:= 0 ; i < len(s); i++ {
		if i != len(s)-1 && s[i] == s[i+1] {
			cont++
		} else {
			res += string(cont + '0') + string(s[i])
			cont = 1
		}
	}
	return res
}