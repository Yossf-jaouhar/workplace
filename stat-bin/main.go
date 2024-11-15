package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	
	sc := bufio.NewScanner(file)
	nms := make([]int, 0) 

	
	for sc.Scan() {
		
		nm, err := strconv.Atoi(sc.Text())
		if err != nil {
			fmt.Println("Invalid number in file, skipping:", sc.Text())
			continue 
		}
		nms = append(nms, nm) 
	}

	if err := sc.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	Average(nms)

	vv := Median(nms)
    fmt.Printf("Median: %2f\n", vv)

	Variance(nms)
}


func Average(nms []int){
	if len(nms) == 0 {
		fmt.Println("No numbers to calculate average.")
		return
	}

	sum := 0
	for _, num := range nms {
		sum += num
	}

	average := float64(sum) / float64(len(nms))

	fmt.Printf("Average: %.2f\n", average)
}



func Median(nms []int) float64{
	md := len(nms)/2

	if md % 2 == 0 {
		mdd := nms[md] + nms[md]+1
		nmm := float64(mdd)
		return nmm
		
	} else {
		nmm := float64(nms[md])
		
		return nmm
	}

}


func Variance(nms []int) {
	md := Median(nms)

	lo := float64(nms[0])
	a := float64(nms[len(nms)-1])

	bn := md - lo
	bnn := md - a


	kk := (bn  + )
}