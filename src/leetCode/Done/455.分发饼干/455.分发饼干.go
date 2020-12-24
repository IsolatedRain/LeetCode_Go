package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	gID, sID, gEnd, sEnd := 0, 0, len(g), len(s)
	for gID < gEnd && sID < sEnd {
		if s[sID] >= g[gID] {
			gID++
		}
		sID++

	}
	return gID
}

func main() {
	g := []int{1, 2}
	s := []int{2, 3, 1}
	r := findContentChildren(g, s)
	fmt.Println(r)
}
