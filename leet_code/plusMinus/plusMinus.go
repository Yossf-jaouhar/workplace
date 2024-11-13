package main

import "fmt"


func main(){
	arr := []int32{1,2,-1,-2,0,0,5}
	fmt.Println(plusMinus(arr))

}


func plusMinus(arr []int32) (float32, float32, float32){
    s  := 0
	m  := 0
	z  := 0


	zise := len(arr)

	for _ , nm := range arr {
		if nm >0 {
			m++
		}
		if nm <0 {
			s++
		}
		if nm == 0 {
			z++
		}
	}

	n := float32(m) / float32(zise)
    h := float32(s) / float32(zise)
	r := float32(z) / float32(zise)

	return n , h , r
}
