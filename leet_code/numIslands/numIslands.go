package main

import "fmt"

func main() {
	// Example grid
	grid := [][]int{
		{1, 1, 0, 0},
		{1, 0, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 1, 0},
	}

	// Call the function to count the connected components
	result, v := numIslands(grid)
	for _, char := range v {
		fmt.Println(char)
	}
	fmt.Println("Number of connected components:", result)
}

func numIslands(grid [][]int) (int, [][]int) {
	cont := 0
	v := grid
	for i := 0; i < len(v); i++ {
		for j := 0; j < len(v[i]); j++ {
			if v[i][j] == 1 {
				v = dfs(v, i, j)
				cont++
			}
		}
	}
	return cont, v
}

func dfs(grid [][]int, i int, j int) [][]int {
	
	return grid
}
