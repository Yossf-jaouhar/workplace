package main

import (
	"fmt"
	"os"

	"lemin/functions"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	info := functions.Info{}

	err := info.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	
	
	
	info.Res = append(info.Res, info.Tunnels[info.End]...)
	
	
	for _, v := range info.Tunnels[info.End] {
		info.Bfs(v)
		if len(info.UniquePaths) >= 2 {
			info.FindGroups()
		}
	}

	for _, p := range info.UniquePaths {
		fmt.Println(p)
	}
}
