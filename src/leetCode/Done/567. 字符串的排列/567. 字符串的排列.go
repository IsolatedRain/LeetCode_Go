package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	count := map[byte]int{}
	for i := range s1 {
		count[s1[i]]++
	}
	L := 0
	for R := range s2 {
		count[s2[R]]--
		for count[s2[R]] < 0 {
			count[s2[L]]++
			L++
		}
		if R-L+1 == len(s1) {
			return true
		}
	}
	return false
}

func main() {
	s1 := "ab"
	s2 := "eidbcaooo"
	fmt.Println(checkInclusion(s1, s2))
	// fmt.Println(checkInclusion("adc", "dcda"))
}
