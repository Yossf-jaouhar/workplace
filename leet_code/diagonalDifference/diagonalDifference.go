package main

import "fmt"

func main() {
	arr := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}

	result := diagonalDifference(arr)
	fmt.Println(result)
}

func diagonalDifference(arr [][]int32) int32 {
	
	var a int32
	var b int32

	for i:=0; i < len(arr); i++ {
		a += arr[i][i]
	

	}

	l := len(arr)
	for i:=0; i < len(arr); i++ {
		b += arr[i][l-1]

	

		l--
	}

	v := a - b
	return v
}
