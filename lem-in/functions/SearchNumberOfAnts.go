package functions

import (
	"strconv"
	"strings"
)

func (a *Info) SearchNumberOfAnts(lines []string) (string , int){
	status := false
	for i:=0; i < len(lines); i++ {
		status = NumberOfAnts(lines[i])
		if status {
			a.NumberOfAnts = strings.TrimSpace(lines[i])
			return "" , i
		} else {
			continue
		}
	}
	v, _ := strconv.Atoi(a.NumberOfAnts)
	if v <= 0 {
		return "Invalid number of ants" , 0
	}
	return "Invalid number of ants" , 0
}

func NumberOfAnts(line string) bool {
	if len(line) == 0 {
		return false
	}
	
	for _, char := range line {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
