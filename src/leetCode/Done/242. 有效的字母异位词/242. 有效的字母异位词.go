package main

import "fmt"

func isAnagram(s string, t string) bool {
	sCount := map[byte]int{}
	tCount := map[byte]int{}
	for i := range s {
		sCount[s[i]]++
	}
	for i := range t {
		tCount[t[i]]++
	}
	for k, v := range sCount {
		tVal, ok := tCount[k]
		if !ok || tVal != v {
			return false
		}
	}
	return true
}

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
}
