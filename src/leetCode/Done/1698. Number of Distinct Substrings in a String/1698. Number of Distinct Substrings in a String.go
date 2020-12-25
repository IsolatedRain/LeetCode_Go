package main

import "fmt"

func countDistinct(s string) int {
	n := len(s)
	mapSet := map[string]struct{}{}
	for i := 0; i < n; i++ {
		for j := n; j > i; j-- {
			if _, ok := mapSet[s[i:j]]; ok {
				break
			}
			mapSet[s[i:j]] = struct{}{}
		}
	}
	return len(mapSet)
}

func main() {
	s := "aabbaba"
	r := countDistinct(s)
	fmt.Println(r)
}
