package functions

import "fmt"

func (a *Info) AllPath() {
	//visited := make(map[string]bool)
	//var allPaths [][]string
	//var path []string
	//Dsf(a.Tunnels, &path, visited, a.Start, a.End, &allPaths)
	allPaths := Bfs(a.Tunnels, a.Start , a.End)
	a.AllPaths = allPaths
}



func Bfs(graph map[string][]string , start string , end string) [][]string {
	var queue [][]string 
	var paths [][]string

	queue = append(queue, []string{start})

	for len(queue) > 0 {
		fmt.Println(queue)
		path := queue[0]
		queue = queue[1:]

		lastroom := path[len(path)-1]
		if lastroom == end {
			paths = append(paths, path)
			continue
		}

		for _, nei := range graph[lastroom] {
			if !isvesited(path , nei) {
				newpath := path
				newpath = append(newpath, nei)
				queue = append(queue, newpath)
			}
		}
	}
	return paths
}

func isvesited(path []string , room string) bool {
	for _, char := range path {
		if char == room {
			return true
		}
	}
	return false
}





func Dsf(graph map[string][]string, path *[]string, visited map[string]bool, start, end string, allPaths *[][]string) {
	*path = append(*path, start)
	visited[start] = true
	if start == end {
		*allPaths = append(*allPaths, append([]string{}, *path...))
	} else {
		for _, neighbor := range graph[start] {
			if !visited[neighbor] {
				Dsf(graph, path, visited, neighbor, end, allPaths)
			}
		}
	}
	*path = (*path)[:len(*path)-1]
	visited[start] = false
}
