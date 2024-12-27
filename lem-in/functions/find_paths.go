package functions

func (y *Info) Bfs(n string) {
	var result [][]string
	var queue [][]string

	queue = append(queue, []string{n})

	for len(queue) > 0 {

		path := queue[0]

		queue = queue[1:]

		lastroom := path[len(path)-1]

		if lastroom == y.End {

			newpath := append([]string{}, path...)

			result = append(result, newpath)
			y.AllGroups = append(y.AllGroups, result)
			break
		}

		for _, nei := range y.Tunnels[lastroom] {
			if !isvesited(path, nei) && nei != y.Start && ok(y.Res, nei, y.End) {
				newpath := append([]string{}, path...)
				newpath = append(newpath, nei)
				queue = append(queue, newpath)
			}
		}
	}
}

func (y *Info) FindMorePaths() {
	if len(y.AllGroups) == 0 {
		return
	}

	for i := 0; i < len(y.AllGroups); i++ {

		newGroup := y.BBfs(y.AllGroups[i][0])
		y.UniqueGroups = append(y.UniqueGroups, newGroup)

	}
}

func (y *Info) BBfs(key []string) [][]string {
	var result [][]string
	result = append(result, key)
	var queue [][]string

	queue = append(queue, []string{y.Start})

	for len(queue) > 0 {

		path := queue[0]

		queue = queue[1:]

		lastroom := path[len(path)-1]

		if len(y.Tunnels[lastroom]) > 2 {
			continue
		}

		if lastroom == y.End {
			newpath := append([]string{}, path...)

			if ValidPath(key, newpath[1:]) {
				
				result = append(result, newpath[1:])
				
			}
			continue
		}

		for _, nei := range y.Tunnels[lastroom] {
			if !isvesited(path, nei) && nei != y.Start {

				newpath := append([]string{}, path...)
				newpath = append(newpath, nei)
				queue = append(queue, newpath)

			}
		}
	}
	return result
}

func ok(n []string, a string, end string) bool {
	if len(n) == 0 {
		return true
	}
	for _, char := range n {
		if char == a && char != end {
			return false
		}
	}
	return true
}

func isvesited(path []string, room string) bool {
	for _, char := range path {
		if char == room {
			return true
		}
	}
	return false
}

func ValidPath(p1 []string, p2 []string) bool {
	if len(p1) == 1 && len(p2) == 1 {
		return false
	}
	
	for i := 0; i < len(p1)-1; i++ {
		for j := 0; j < len(p2)-1; j++ {
			if p1[i] == p2[j] {
				return false
			}
		}
	}
	return true
}

