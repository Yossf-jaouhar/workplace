package functions

import (
	"fmt"
	"strconv"
)

func (y *Info) Print(res [][]string) {
	result := Sort(res)

	NubOfAnts := conv(y.NumberOfAnts)

	var steps [][]string
	var step []string

	for len(NubOfAnts) > 0 {
		for i := 0; i < len(result); i++ {
			if i != len(result)-1 && len(result[i]) <= len(result[i+1]) {

				result[i] = append(result[i], NubOfAnts[0])

				NubOfAnts = NubOfAnts[1:]

				step = result[i]

				steps = append(steps, step)

			} else if i != len(result)-1 && len(result[i]) > len(result[i+1]) {

				result[i+1] = append(result[i+1], NubOfAnts[0])

				NubOfAnts = NubOfAnts[1:]

				step = result[i+1]

				steps = append(steps, step)

			}
		}
	}

	y.print2(steps)
}

func (y *Info) print2(res [][]string) {

	var mat [][]string

	for i := len(res) - 1; i >= 0; i-- {
		if len(mat) <= y.NumberOfAnts {
			g := ggg(res[i])
			mat = append(mat, g...)
		} else {
			break
		}
	}

	for _, v := range mat {
		fmt.Println(v)
	}
}

func (y * Info) ggg(res []string) [][]string {
	
	NubOfAnts := conv(y.NumberOfAnts)

	var nml []string

	var nmll [][]string

	hi := false

	for _, v := range res {
		if !hi && v != "end" {
			nml = append(nml, v)
		} else if v == "end" {
			hi = true
			continue
		} else if hi && v != "end" {
			vv := append([]string{}, v)
			vv = append(vv, nml...)
			nmll = append(nmll, vv)
		}
	}
	return nmll
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
