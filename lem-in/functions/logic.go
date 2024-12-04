package functions

import "fmt"

func (a *Info) Logic() {
	var path []string

	vesited := make(map[string]bool)

	res , b := Dsf(a.Tunnels , path, vesited, a.Start , a.End)
	if !b {
		a.AllPaths = append(a.AllPaths, res)
	}
	
	
}

func Dsf(graph map[string][]string , path []string, v map[string]bool, start string , end string) ([]string , bool){

	var all []string

	v[start] = true

	path = append(path, start)

	all = path

	if start == end {
		fullPath := append([]string{}, path...)
        fmt.Println("Path Found:", fullPath)

		return fullPath , true
 	}

	for _, char := range graph[start] {
		if !v[char] {
			res , b := Dsf(graph , path , v , char , end)
			return res , b
		}
	}
	path = path[:len(path)-1]
	return all , false
}
