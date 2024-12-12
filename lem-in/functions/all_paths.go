package functions

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
			newpath := append([]string{}, path...)
			a.UniquePaths = append(a.UniquePaths, newpath)
			continue
		}

		for _, nei := range a.Tunnels[lastroom] {
			if !isvesited(path, nei) {
				a.Neiofstart[nei] = true
			}
		}
	}
	
}

func isvesited(path []string, room string) bool {
	for _, char := range path {
		if char == room {
			return true
		}
	}
	return false
}

