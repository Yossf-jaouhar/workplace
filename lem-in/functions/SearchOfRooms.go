package functions

import (
	"strings"
)

func (a *Info) SearchOfRooms(lines []string, Index int) string {
	lines = lines[Index+1:]
	status := false
	for i := 0; i < len(lines); i++ {
		lines[i] = strings.TrimSpace(lines[i])
		if lines[i] == "##start" {
			status = false
			if i < len(lines)-1 {
				for !status {
					
					Start := i + 1
					if Start == len(lines) -2 {
						return "There is no Start room"
					}
					if a.Validation(lines[Start]) {
						StartRoom := Select_Start_End(lines[Start])
						a.Start = StartRoom
						status = true
						break
					} else {
						Start++
					}
				}
			} else {
				return "There is no Start room"
			}
		} else if lines[i] == "##end" {
			status = false
			if i < len(lines)-1 {
				for !status {
					End := i + 1
					if End == len(lines) -2 {
						return "There is no end room"
					}
					if a.Validation(lines[End]) {
						EndRoom := Select_Start_End(lines[End])

						a.End = EndRoom
						status = true
						break
					} else {
						End++
					}
				}
			}  else {
				return "There is no  End room"
			}
		}
	}

	return ""
}

func Select_Start_End(Room string) string {
	if len(Room) == 0 {
		return ""
	}
	RoomStartOrEnd := strings.Fields(Room)
	if len(RoomStartOrEnd) > 3 || len(RoomStartOrEnd) < 3 {
		return ""
	}
	return RoomStartOrEnd[0]
}

func (a *Info) Validation(line string) bool {
	if len(line) == 0 {
		return false
	}
	Everyting := strings.Fields(line)
	if len(Everyting) > 3 || len(Everyting) < 3 {
		return false
	}

	ms := Check(Everyting[0])
	if ms != "" {
		return false
	}

	if !checkCoordinates(Everyting[1], Everyting[2]) {
		return false
	}

	if a.Rooms == nil {
		a.Rooms = make(map[string][]string)
	}

	a.Rooms[Everyting[0]] = append(a.Rooms[Everyting[0]], Everyting[1])
	a.Rooms[Everyting[0]] = append(a.Rooms[Everyting[0]], Everyting[2])

	return true
}

func Check(str string) string {
	if str[0] == '#' || str[0] == 'L' {
		return "Invalid room"
	}
	return ""
}

func checkCoordinates(t1, t2 string) bool {
	for _, char := range t1 {
		if char < '0' || char > '9' {
			return false
		}
	}
	for _, char := range t2 {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
