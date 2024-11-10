package main

import "fmt"

func main() {
	sl := []string{"flower", "flower", "flo"}
	fmt.Println(longestCommonPrefix(sl))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	
	for _, st := range strs[1:] {
		for len(st) < len(prefix) || st[:len(prefix)] != prefix {
		
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}
