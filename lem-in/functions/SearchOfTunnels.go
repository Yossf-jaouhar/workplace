package functions

import (
	"strings"
)

func (a *Info) SearchOfTunnels(lines []string, index int) string {
	for i := index + 1; i < len(lines); i++ {
		if checkTunnels(lines[i]) {
			Tunnels := strings.Split(lines[i], "-")
			if len(Tunnels) == 2 {
				var room1 string
				var room2 string
				if a.ValidRoom(Tunnels[0], Tunnels[1]) {
					room1 = Tunnels[0]
					room2 = Tunnels[1]
				} else {
					return "Invalid Tunnels"
				}
				a.FillTheGraph(room1, room2)

			}
		}
	}

	return ""
}

func checkTunnels(str string) bool {
	for _, char := range str {
		if char == '-' {
			return true
		}
	}
	return false
}

func (a *Info) FillTheGraph(room1, room2 string) {
	if a.Tunnels == nil {
		a.Tunnels = map[string][]string{}
	}

	a.Tunnels[room1] = append(a.Tunnels[room1], room2)
}

func (a *Info) ValidRoom(room1, room2 string) bool {
	s1 := false
	s2 := false
	for char := range a.Rooms {
		if char == room1 {
			s1 = true
		}
		if char == room2 {
			s2 = true
		}
	}
	if s1 && s2 {
		return true
	} else {
		return false
	}
}
