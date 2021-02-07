package main

import (
	"fmt"
	"sort"
)

func findLongestWord(s string, d []string) string {
	sort.Slice(d, func(i, j int) bool { return len(d[i]) > len(d[j]) || len(d[i]) == len(d[j]) && d[i] < d[j] })
	isValid := func(i int) bool {
		wID, sID := 0, 0
		for wID < len(d[i]) && sID < len(s) {
			if d[i][wID] == s[sID] {
				wID++
			}
			sID++
		}
		return wID == len(d[i])
	}
	for i := range d {
		if isValid(i) {
			return d[i]
		}
	}
	return ""
}

func main() {
	fmt.Println(findLongestWord("abpcplea",
		[]string{"ale", "apple", "monkey", "plea", "abcle"}))
}
