package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() 

	buffer, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}


	fmt.Println(string(buffer))
}
