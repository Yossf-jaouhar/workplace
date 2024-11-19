package main

import "fmt"

func main() {
	s := "abcabcbb"
    ss := "pwwkew"
	u := lengthOfLongestSubstring(s)
    v := lengthOfLongestSubstring(ss)
  
	fmt.Println(u)
    fmt.Println(v)
}

func lengthOfLongestSubstring(s string) int {
	max := 0
	st := 0
	en := 1

	for i:=0 ; i < len(s); i++ {
		v := s[st:en+1]
		if isvalid(v) {
			if max < len(v) {
				max = len(v)
			}
			en++
		} else if !(isvalid(v)) {
		
			st++
		}
	
	}




	return max
	
}


func isvalid(res string) bool {
	mapp := make(map[rune]bool)
	
	for _, char := range res {
		if mapp[char] {
			return false
		}
		mapp[char] = true
	}
	return true
}