package functions

import (
	
	"strings"
)

func (a *Info) SearchOfTunnels(lines []string) string {
	for i := 0; i < len(lines); i++ {
		if checkTunnels(lines[i]) {
			Tunnels := strings.Split(lines[i], "-")
			if len(Tunnels) == 2 {
				room1 := Tunnels[0]
				room2 := Tunnels[1]
				Message := a.FillTheGraph(room1, room2)
				if Message != "" {
					return Message
				}
			} else {
				return "Invalid Tunnels"
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

func (a * Info) FillTheGraph(room1, room2 string) string {
	
	return ""
}