package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"lemin/functions"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		return
	}

	sc := bufio.NewScanner(file)
	var lines []string
	for sc.Scan() {
		l := sc.Text()
		if l == "" {
			continue
		}
		line := strings.TrimSpace(l)

		lines = append(lines, line)
	}
	info := functions.Info{}
	MessageOfInvalidInput := info.SearchNumberOfAntsAndRoomsAndTunnels(lines)
	if MessageOfInvalidInput != nil {
		fmt.Println(MessageOfInvalidInput)
		return
	}

	info.Bfs()

	//fmt.Println(info.NumberOfAnts)
	//fmt.Println(info.Start)
	//fmt.Println(info.End)
	//fmt.Println(info.Rooms)
	//fmt.Println(info.Tunnels)
	for _, p := range info.UniquePaths {
		fmt.Println(p)
	}
}
