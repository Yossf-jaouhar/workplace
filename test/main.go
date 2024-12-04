package main

import "fmt"

func dfs(graph map[string][]string, start string, visited map[string]bool) {
	visited[start] = true
	fmt.Println("Visited:", start)

 	for _, neighbor := range graph[start] {
		if !visited[neighbor] {
			dfs(graph, neighbor, visited)
		}
	}
	
}

func main() {
	graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"D"},
		"C": {"E"},
		"D": {"F"},
		"E": {"F"},
		"F": {},
	}

	visited := make(map[string]bool)

	fmt.Println("Depth-First Search Traversal:")
	dfs(graph, "A", visited)
}
