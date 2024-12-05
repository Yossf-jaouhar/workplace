package functions

func (a *Info) AllPath() {
	visited := make(map[string]bool)
	var allPaths [][]string
	var path []string
	Dsf(a.Tunnels, &path, visited, a.Start, a.End, &allPaths)

	a.AllPaths = allPaths
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
