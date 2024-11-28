package functions

import (
	"fmt"
	"strings"
)

func (a *Info) SearchNumberOfAntsAndRoomsAndTunnels(lines []string) string {
	var MessageOfInvalidInput string

	MessageOfInvalidInput = a.SearchNumberOfAnts(lines)

	if MessageOfInvalidInput != "" {
		return MessageOfInvalidInput
	}

	MessageOfInvalidInput = a.SearchOfRooms(lines)

	if MessageOfInvalidInput != "" {
		return MessageOfInvalidInput
	}
	return ""
}

func (a *Info) SearchNumberOfAnts(lines []string) string {
	Number := lines[0]
	var NmAnts string

	for _, char := range Number {
		if char < '0' || char > '9' {
			return "invalid number of ants"
		} else {
			NmAnts += string(char)
		}
	}
	a.NumberOfAnts = NmAnts

	return ""
}

func (a *Info) SearchOfRooms(lines []string) string {
	lines = lines[1:]
	var RoomsAndStartAndEnd []string
	Start := false
	End := false
	for _, line := range lines {
		if line == "##end" && !End {
			End = true
			continue
		}
		if line == "##start" || Start {
			if Start {
				RoomsAndStartAndEnd = append(RoomsAndStartAndEnd, line)
			}
			Start = true
			if End {
				End = false
				Start = false
			}
		}
	}
	a.SelectRooms(RoomsAndStartAndEnd)
	return ""
}

func (a *Info) SelectRooms(RoomsAndStartAndEnd []string) {
	RSE := RoomsAndStartAndEnd
	var res []string
	for i := 0; i < len(RSE); i++ {
		if i == 0 {
			if Validation(string(RSE[i])) {
				a.Start = string(RSE[i][0])
				res = append(res, RSE[i])
			}
		} else if i == len(RSE)-1 {
			if Validation(string(RSE[i])) {
				a.End = string(RSE[i][0])
				res = append(res, RSE[i])
			}
		} else if Validation(string(RSE[i])) {
			res = append(res, RSE[i])
		} else {
			continue
		}
	}
	fmt.Println(res)
	a.FillTheMapOfRooms(res)
}

// Validation Of Rooms And Start And End And Select everyting.
func Validation(Everyting string) bool {
	FillTheRoom := false
	Room := ""

	FillTheX := false
	x := ""

	FillTheY := false
	y := ""

	for i := 0; i < len(Everyting); i++ {
		if Everyting[i] != ' ' && !FillTheRoom && !FillTheX && !FillTheY {
			if Everyting[i] >= '0' && Everyting[i] <= '9' {
				Room += string(Everyting[i])
			} else {
				return false
			}
			if i != len(Everyting)-1 && Everyting[i+1] == ' ' {
				FillTheRoom = true
			}
			continue
		} else if Everyting[i] != ' ' && FillTheRoom && !FillTheX && !FillTheY {
			if Everyting[i] >= '0' && Everyting[i] <= '9' {
				x += string(Everyting[i])
			} else {
				return false
			}

			if i != len(Everyting)-1 && Everyting[i+1] == ' ' {
				FillTheX = true
			}

			continue
		} else if Everyting[i] != ' ' && FillTheRoom && FillTheX && !FillTheY {
			if Everyting[i] >= '0' && Everyting[i] <= '9' {
				y += string(Everyting[i])
			} else {
				return false
			}

			if i != len(Everyting)-1 && Everyting[i+1] == ' ' {
				FillTheY = true
			}

			continue
		}
	}

	return true
}

func (a *Info) FillTheMapOfRooms(rooms []string) {
	b := false
	c := 0
	for _, room := range rooms {
		if 
	}
}
