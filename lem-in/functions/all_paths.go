package functions

// func (a *Info) AllPath() {
// 	// visited := make(map[string]bool)
// 	// var allPaths [][]string
// 	// var path []string
// 	// Dsf(a.Tunnels, &path, visited, a.Start, a.End, &allPaths)
//     a.Bfs(a.Tunnels, a.Start, a.End)
// 	// a.AllPaths = append(a.AllPaths, allPath...)
// 	//a.AllPaths = allPaths
// }

func (a *Info) Bfs() {
	if a.Neiofstart == nil {
		a.Neiofstart = make(map[string]bool)
	}

	var queue [][]string

	queue = append(queue, []string{a.Start})

	for len(queue) > 0 {

		path := queue[0]

		queue = queue[1:]

		lastroom := path[len(path)-1]

		if lastroom == a.End {

			
			continue
		}

		for _, nei := range a.Tunnels[lastroom] {
			if !isvesited(path, nei) {
				a.Neiofstart[nei] = true
			}
		}
	}
	
}

// func (a *Info) FindTheBestPaths() {
// 	status := false

// 	sma := 0
// 	bg := 0
	
// 	for j := 0 ; j < len(a.AllPaths); j++ {
// 		for i := 1; i < len(a.AllPaths)-1; i++ {

// 			for k := 0; k < len(a.AllPaths[i]); k++ {

// 				if i < len(a.AllPaths)-1 && j < len(a.AllPaths)-1 {
// 					status = true
// 				}
// 				if !status {
// 					a.UniquePaths = append(a.UniquePaths, p)
// 				}

// 			}

// 		}
// 	}
// }

func isvesited(path []string, room string) bool {
	for _, char := range path {
		if char == room {
			return true
		}
	}
	return false
}

// func Dsf(graph map[string][]string, path *[]string, visited map[string]bool, start, end string, allPaths *[][]string) {
// 	*path = append(*path, start)
// 	visited[start] = true
// 	if start == end {
// 		*allPaths = append(*allPaths, append([]string{}, *path...))
// 	} else {
// 		for _, neighbor := range graph[start] {
// 			if !visited[neighbor] {
// 				Dsf(graph, path, visited, neighbor, end, allPaths)
// 			}
// 		}
// 	}
// 	*path = (*path)[:len(*path)-1]
// 	visited[start] = false
// }
