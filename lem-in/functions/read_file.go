package functions

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

func (a *Info) ReadFile(arg string) ([]string, error) {
	file, err := os.Open(arg)
	if err != nil {
		return nil, errors.New("error to open the file")
	}

	lines := make([]string, 20)

	sc := bufio.NewScanner(file)

	na := false
	for sc.Scan() {
		l := sc.Text()
		if l == "" {
			continue
		}
		line := strings.TrimSpace(l)
		if !na {
			var err error

			a.NumberOfAnts, err = strconv.Atoi(lines[0])
			if err != nil {
				return nil, errors.New("invalid nm ants")
			}
			if a.NumberOfAnts <= 0 {
				return nil, errors.New("invalid nm ants")
			}
		}
		
	}

	return lines, nil
}
