package main

import "fmt"


func main(){
	res := 0
    result := ""
	for i:=9;i>=1; i--{
		for j := 9; j >=1; j--{
			for  k := 9; k>=1; k--{
				if i > j && i > k && j > k{
					res = i*100+j*10+k

					result += changeg(res)
					result += ", "
				}
			}
		}
	}
	if len(result) > 2 {
		result = result[:len(result)-2]
	}
	fmt.Println(result)
}

func changeg(nm int) string{
	res := ""
	for nm != 0 {
		r := nm % 10
		res = string(r+'0') + res
		nm = nm /10
	}
	return res
}