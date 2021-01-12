package main

import (
	"fmt"
	"sort"
)

func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}
	var getRoot func(int) int
	getRoot = func(x int) int {
		if p[x] != p[p[x]] {
			p[x] = getRoot(p[x])
		}
		return p[x]
	}
	var union func(int, int)
	union = func(x, y int) {
		xRoot, yRoot := getRoot(x), getRoot(y)
		if xRoot > yRoot {
			p[xRoot] = yRoot
		} else {
			p[yRoot] = xRoot
		}
	}
	for _, v := range pairs {
		union(v[0], v[1])
	}
	chars := map[int][]byte{}
	for i := range p {
		k := getRoot(i)
		chars[k] = append(chars[k], s[i])
	}

	for k := range chars {
		sort.Slice(chars[k], func(i, j int) bool { return chars[k][i] < chars[k][j] })
	}

	res := make([]byte, n)
	for i := range res {
		res[i] = chars[p[i]][0]
		chars[p[i]] = chars[p[i]][1:]
	}
	return string(res)
}

func main() {
	s := "dcab"
	pairs := [][]int{{0, 3}, {1, 2}}
	fmt.Println(smallestStringWithSwaps(s, pairs))
}
