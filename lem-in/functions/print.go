package functions

import (
	"fmt"
	"strconv"
)

func (y *Info) Print(res [][]string) {
	res = Sort(res)
	NubOfAnts := conv(y.NumberOfAnts)

	var steps [][]string
	var step []string

	result := res

	fmt.Println(NubOfAnts)

	for len(steps) != len(NubOfAnts)  {
	
		for i := 0; i < len(result); i++ {

			if i != len(result)-1 && len(result[i]) <= len(result[i+1]) {

			} else if i != len(result)-1 && len(result[i]) > len(result[i+1]) {
				
			}
		}
	}

	y.hhhhhhhh(steps)
	
}







func (y *Info) hhhhhhhh(res [][]string) {
	for _, v := range res {
		fmt.Println(v)
	}
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
		if i != len(res)-1 && len(res[i]) > len(res[i+1]) {
			res[i], res[i+1] = res[i+1], res[i]
		} else {
			continue
		}
	}
	return res
}
