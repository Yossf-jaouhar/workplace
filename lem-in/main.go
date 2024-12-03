package main

import (
	"bufio"
	"fmt"
	"os"

	"lemin/functions"
)

func main() {
	file, err := os.Open("text.txt")
	if err != nil {
		return
	}

	sc := bufio.NewScanner(file)
	var lines []string
	for sc.Scan() {
		line := sc.Text()
		lines = append(lines, line)
	}
	info := functions.Info{}
	MessageOfInvalidInput := info.SearchNumberOfAntsAndRoomsAndTunnels(lines)
	if MessageOfInvalidInput != "" {
		fmt.Println(MessageOfInvalidInput)
		return
	}

	info.Logic()

	// fmt.Println(info.NumberOfAnts)
	// fmt.Println(info.Start)
	// fmt.Println(info.End)
	// fmt.Println(info.Rooms)
	// fmt.Println(info.Tunnels)
}
