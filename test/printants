package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func generateLargeLemInFile(fileName string, numRooms, numLinks, numAnts int) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	rand.Seed(time.Now().UnixNano())

	_, err = file.WriteString(fmt.Sprintf("%d\n", numAnts))
	if err != nil {
		return fmt.Errorf("error writing number of ants: %w", err)
	}

	_, err = file.WriteString("##start\n")
	if err != nil {
		return fmt.Errorf("error writing start room: %w", err)
	}
	_, err = file.WriteString("0 0 0\n")
	if err != nil {
		return fmt.Errorf("error writing start room coordinates: %w", err)
	}

	_, err = file.WriteString("##end\n")
	if err != nil {
		return fmt.Errorf("error writing end room: %w", err)
	}
	endRoomCoordinates := fmt.Sprintf("%d %d %d\n", numRooms-1, rand.Intn(1000), rand.Intn(1000))
	_, err = file.WriteString(endRoomCoordinates)
	if err != nil {
		return fmt.Errorf("error writing end room coordinates: %w", err)
	}

	for i := 1; i < numRooms-1; i++ {
		roomCoordinates := fmt.Sprintf("%d %d %d\n", i, rand.Intn(1000), rand.Intn(1000))
		_, err = file.WriteString(roomCoordinates)
		if err != nil {
			return fmt.Errorf("error writing room %d: %w", i, err)
		}
	}

	links := make(map[string]struct{})
	for i := 0; i < numLinks; i++ {
		var roomA, roomB int
		for attempts := 0; attempts < 100; attempts++ { 
			roomA = rand.Intn(numRooms)
			roomB = rand.Intn(numRooms)
			if roomA != roomB {
				link := fmt.Sprintf("%d-%d", roomA, roomB)
				reverseLink := fmt.Sprintf("%d-%d", roomB, roomA)
				if _, exists := links[link]; !exists {
					if _, exists := links[reverseLink]; !exists {
						links[link] = struct{}{}
						break
					}
				}
			}
		}
		_, err = file.WriteString(fmt.Sprintf("%d-%d\n", roomA, roomB))
		if err != nil {
			return fmt.Errorf("error writing link %d-%d: %w", roomA, roomB, err)
		}
	}

	return nil
}

func main() {

	numRooms := 1000000  
	numLinks := 5000000  
	numAnts := 10000000  

	err := generateLargeLemInFile("large_lem_in_file.txt", numRooms, numLinks, numAnts)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Large Lem-in file generated successfully!")
	}
}
