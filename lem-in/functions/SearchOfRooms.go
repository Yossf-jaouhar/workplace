package functions

import (
	"errors"
	"strings"
)

func (a *Info) SearchOfRooms(lines []string) error {
	lines = lines[1:]
	StatusStart := false
	StatusEnd := false
	
	for i := 0; i < len(lines); i++ {

		lines[i] = strings.TrimSpace(lines[i])
		if a.Validation(lines[i]) {
			if i > 0 {
				g := i-1	
				for g >= 0{
					if !StatusStart && lines[g] == "##start" {
						res := strings.Fields(lines[i])
						a.Start = res[0]
						StatusStart = true
						break
					}
					if !StatusEnd && lines[g] == "##end" {
						res := strings.Fields(lines[i])
						a.End = res[0]
						StatusEnd = true
						break	
					}
					g--
				}
				
			}
		}
	}
	if !Vld(lines) {
		return errors.New("invali form of the file")
	}
	if !StatusStart {
		return errors.New("no start room")
	}
	if !StatusEnd {
		return errors.New("no end room")
	}
	
	return nil
}


func Vld(lines []string) bool {
	s, e := 0, 0
	for _, char := range lines {
		if char == "##start" {
			s++
		} else if char == "##end" {
			e++
		}
	}
	if e == 1 && s == 1 {
		return true
	}
	return false
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
