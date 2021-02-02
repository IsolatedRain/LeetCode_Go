package main

import "fmt"

func numSimilarGroups(strs []string) int {
	n := len(strs)
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}
	var getRoot func(x int) int
	getRoot = func(x int) int {
		if p[x] != p[p[x]] {
			p[x] = getRoot(p[x])
		}
		return p[x]
	}
	var union func(x, y int)
	union = func(x, y int) {
		xRoot, yRoot := getRoot(x), getRoot(y)
		if xRoot != yRoot {
			p[yRoot] = xRoot
		}
	}
	var isSimilar func(a, b string) bool
	isSimilar = func(a, b string) bool {
		diff := 0
		for i := range a {
			if a[i] != b[i] {
				diff++
			}
			if diff > 2 {
				return false
			}
		}
		return true
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isSimilar(strs[i], strs[j]) {
				union(i, j)
			}
		}
	}

	res := 0
	for i := range p {
		if p[i] == i {
			res++
		}
	}
	return res
}

func main() {
	strs := []string{"tars", "rats", "arts", "star"}
	fmt.Println(numSimilarGroups(strs))
}
