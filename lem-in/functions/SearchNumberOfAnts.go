package functions

import (
	"strconv"
)

func (a *Info) SearchNumberOfAnts(lines []string) string {
	v, err := strconv.Atoi(lines[0])
	if err != nil {
		return "invalid number of ants"
	}

	if v <= 0 {
		return "Invalid number of ants"
	}
	a.NumberOfAnts = v
	return ""
}
