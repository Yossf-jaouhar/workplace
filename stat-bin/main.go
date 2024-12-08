package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"hh/functions"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	var nms []int

	for sc.Scan() {
		nm, err := strconv.Atoi(sc.Text())
		if err != nil {
			fmt.Println("Invalid number in file", sc.Text())
			return
		}
		nms = append(nms, nm)
	}

	if err := sc.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	average := functions.Average(nms)
	median := functions.Median(nms)
	variance := functions.Variance(nms, average)
	//stdDev := functions.StandardDeviation(variance)

	fmt.Println("Average:", average)
	fmt.Println("Median:", median)
	fmt.Println("Variance:", variance)
	//fmt.Println("Standard Deviation:", stdDev)
}
