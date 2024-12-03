package functions

import "fmt"

var path []string

func (a *Info) Logic() {

	viseted := make(map[string]bool)
	
	for r := range a.Tunnels {
		if r == a.Start {
			dfs(a.Tunnels, r, viseted, path )
		}
	}

	fmt.Println(path)
}

func dfs(mapp map[string][]string, start string , v map[string]bool , path []string ) {

	

	if v[start] {
		return
	}

	v[start] = true

	for char, tt := range mapp {
		for _, dd := range tt {

			if char == start {
				dfs(mapp, dd, v, path)
			}
		}
		path = append(path, char)
	}
}
