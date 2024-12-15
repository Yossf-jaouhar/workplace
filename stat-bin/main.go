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
	var nms []float64

	for sc.Scan() {
		if len(sc.Text()) == 0 {
			continue
		}
		nm, err := strconv.Atoi(sc.Text())
		if err != nil {
			fmt.Println("Invalid number in file", sc.Text())
			return
		}
		nms = append(nms, float64(nm))
	}

	if err := sc.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	average := functions.Average(nms)
	median := functions.Median(nms)
	variance := functions.Variance(nms, average)
	stdDev := functions.StandardDeviation(variance)

	fmt.Printf("Average: %.0f\n", average)
	fmt.Printf("Median: %.0f\n", median)
	fmt.Printf("Variance: %.0f\n", variance)
	fmt.Printf("Standard Deviation: %.0f\n", stdDev)
}
