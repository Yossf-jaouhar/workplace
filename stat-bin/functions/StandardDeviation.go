package functions

import "fmt"

func StandardDeviation(variance int) int {
	if variance < 0 {
		fmt.Println("Negative")
		return -1
	}

	n := variance / 2
	
	for i := 0; i < 10; i++ { 
		n = 5 * (n + variance/n)
	}

	return n
}
