package functions

import (
	"fmt"
	"strconv"
	"strings"
)

func (y *Info) Print(res [][]string) {
	result := Sort(res)
	var newresult [][]string
	newresult = append(newresult, result...)
	NubOfAnts := conv(y.NumberOfAnts)

	var steps [][]string
	var step []string

	for len(NubOfAnts) > 0 {
		for i := 0; i < len(result); i++ {
			if i != len(result)-1 && len(result[i]) <= len(result[i+1]) {

				result[i] = append(result[i], NubOfAnts[0])

				step = append(step, NubOfAnts[0])

				step = append(step, newresult[i]...)

				steps = append(steps, step)

				step = nil

				NubOfAnts = NubOfAnts[1:]

			} else if i != len(result)-1 && len(result[i]) > len(result[i+1]) {

				result[i+1] = append(result[i+1], NubOfAnts[0])

				step = append(step, NubOfAnts[0])

				step = append(step, newresult[i+1]...)

				steps = append(steps, step)

				step = nil

				NubOfAnts = NubOfAnts[1:]

			}
		}
	}

	y.print2(steps)


	y.print3()
}


func (y * Info) print3() {

	var result []string

	var s  string

	

	for _, v := range y.Forprint {

		if s == "" {
			s += v
			continue
		} 

		g := fff(v)

		if !strings.Contains(s , g) {
			s += v
			
		} else {
			result = append(result, s)
			s = v
		}


	}
	var ff []string

	for _, v := range result {
		for _, b := range v {
			
		}	
	}


}

func fff(s string) string {
	parts := strings.Split(s, "-")

	return parts[1]
}

func (y *Info) print2(res [][]string) {
	var arr []string
	var line string

	if res == nil {
		return
	}

	for i := 0; i < len(res); i++ {
		if len(res[i]) <= 1 {
			return
		}
		line = res[i][0] + "-" + res[i][1] + " "

		if strings.Contains(line, res[i][1]) {
			arr = append(arr, line)
			line = ""
		}

		g := res[i][0]

		res[i] = res[i][2:]

		h := append([]string{}, g)

		h = append(h, res[i]...)


		res[i] = h
	}

	for _, v := range arr {
		y.Forprint = append(y.Forprint, v)
	}
	y.print2(res)
}

func conv(nm int) []string {
	var res []string

	for i := 1; i <= nm; i++ {
		l := strconv.Itoa(i)
		res = append(res, "L"+l)

	}
	return res
}

func Sort(res [][]string) [][]string {
	if len(res) == 1 {
		return res
	}

	for i := 0; i < len(res); i++ {
		if i != len(res)-1 && len(res[i]) >= len(res[i+1]) {
			res[i], res[i+1] = res[i+1], res[i]
		} else {
			continue
		}
	}
	return res
}
