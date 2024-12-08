package main

import (
	"fmt"
	"time"
)

// Define a structure to represent an ant with its information:
// - Ant ID (id)
// - Path it follows (path)
// - Current position in the path (position)
type Ant struct {
	id       int
	path     []int
	position int
}

func main() {
	// Define the best paths for the ants
	// Each path is a series of rooms starting from room 0 and ending at room 1
	paths := [][]int{
		{0, 4, 2, 1},       // First path
		{0, 6, 5, 2, 1},    // Second path
		{0, 6, 7, 2, 1},    // Third path
	}

	// Number of ants
	numAnts := 3

	// Create ants and assign them to the paths
	// Each ant is assigned a path based on its order
	ants := []Ant{}
	for i := 1; i <= numAnts; i++ {
		ants = append(ants, Ant{
			id:       i,                          // Ant ID
			path:     paths[(i-1)%len(paths)],   // Assigned path for the ant
			position: 0,                         // Each ant starts at the first room in its path
		})
	}

	// Move the ants step by step
	step := 1 // Step counter
	for {
		fmt.Printf("Step %d:\n", step)
		allReached := true // Variable to check if all ants have reached the end

		// Move each ant
		for i := range ants {
			ant := &ants[i] // Reference to the current ant
			if ant.position < len(ant.path)-1 { // If the ant hasn't reached the final room
				allReached = false             // There's still an ant that hasn't reached the end
				ant.position++                 // Move the ant to the next room
				fmt.Printf("Ant %d moves to room %d\n", ant.id, ant.path[ant.position])
			}
		}

		// Check if all ants have reached the final room, then exit the loop
		if allReached {
			fmt.Println("All ants have reached the final room!")
			break
		}

		// Wait for a second to visualize the movement
		time.Sleep(1 * time.Second)
		step++
	}
}
