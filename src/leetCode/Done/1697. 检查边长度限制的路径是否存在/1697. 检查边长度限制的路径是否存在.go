package main

import (
	"fmt"
	"sort"
)

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
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
	sort.Slice(edgeList, func(i, j int) bool { return edgeList[i][2] < edgeList[j][2] })
	for i := range queries {
		queries[i] = append(queries[i], i)
	}
	sort.Slice(queries, func(i, j int) bool { return queries[i][2] < queries[j][2] })
	res := make([]bool, len(queries))
	i := 0
	for _, q := range queries {
		curLimit := q[2]
		for i < len(edgeList) && edgeList[i][2] < curLimit {
			union(edgeList[i][0], edgeList[i][1])
			i++
		}
		res[q[3]] = false
		if getRoot(q[0]) == getRoot(q[1]) {
			res[q[3]] = true
		}
	}
	return res
}

func main() {
	n := 3
	edgeList := [][]int{{0, 1, 2}, {1, 2, 4}, {2, 0, 8}, {1, 0, 16}}
	queries := [][]int{{0, 1, 2}, {0, 2, 5}}
	fmt.Println(distanceLimitedPathsExist(n, edgeList, queries))
}
