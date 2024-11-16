package main

import (
    "fmt"
    
)

func main(){
    a := []string{"coding", "algorithm", "ascii", "package", "golang"}
    fmt.Printf("%#v\n", Slice(a, 1))
    fmt.Printf("%#v\n", Slice(a, 2, 4))
    fmt.Printf("%#v\n", Slice(a, -3))
    fmt.Printf("%#v\n", Slice(a, -2, -1))
    fmt.Printf("%#v\n", Slice(a, 2, 0))
}


func Slice(a []string, nbrs... int) []string{
	res :=  make([]string,0)

	if len(nbrs) == 1 && nbrs[0] > 0{
		res = a[nbrs[0]:]
		return res 
	} else if len(nbrs) == 1 && nbrs[0] < 0 {
		gg := -nbrs[0]
		res = a[len(a)-gg:]
		return res
	} else if len(nbrs) == 2 && nbrs[0] > 0 && nbrs[1] > 0 {
		res = a[nbrs[0]:nbrs[1]]
		return res
	} else if len(nbrs) == 2 && nbrs[0] < 0 && nbrs[1] < 0 {
		b:= -nbrs[0]
		n := -nbrs[1] 
		for i := len(a)-1; i >= 0; i-- {
			res =append(res, a[i])
		}
		res = res[n:b]
		return res
	} 
 	return nil
}