package functions

import "fmt"

func (a *Info) LogicAnts() {
	var res [][]string
	for _, path := range a.AllPaths {
		res = append(res, path[1:len(path)-1])
	}
	result := UniquePath(res)
	fmt.Println(result)
}

func UniquePath(res [][]string) [][]string {
	var result [][]string
	h := false
	for i := 0; i < len(res); i++ {
		if i < len(res)-1 {
			g := i 
			for g < len(res)-1{
				
				for _, char := range res[g]{
					for _, st := range res[g+1] {
						if char == st {
							h = true	
						}
					}
				}
				g++
			}
		}

		if !h {
			result = append(result,res[i])
		}
	}
	return result
}
