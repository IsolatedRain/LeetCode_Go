package main

import (
	"fmt"
	"sort"
)

type uf struct {
	p    []int
	rank []int
	size int
}

func newUf(n int) *uf {
	arr := make([]int, n)
	r := make([]int, n)
	for i := range arr {
		arr[i] = i
		r[i] = 1
	}

	return &uf{p: arr, rank: r, size: n}
}

func (u *uf) getRoot(x int) int {
	if u.p[x] != u.p[u.p[x]] {
		u.p[x] = u.getRoot(u.p[x])
	}
	return u.p[x]
}
func (u *uf) union(x, y int) bool {
	xRoot := u.getRoot(x)
	yRoot := u.getRoot(y)
	if xRoot == yRoot {
		return false
	}
	if u.rank[xRoot] < u.rank[yRoot] {
		xRoot, yRoot = yRoot, xRoot
	}
	u.p[yRoot] = xRoot
	u.rank[xRoot] += u.rank[yRoot]
	u.size--
	return true
}

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	edgesLen := len(edges)
	for i := range edges {
		edges[i] = append(edges[i], i)
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i][2] < edges[j][2] })
	ufSTD := newUf(n)
	minVal := 0
	for i := 0; i < edgesLen; i++ {
		if ufSTD.union(edges[i][0], edges[i][1]) {
			minVal += edges[i][2]
		}
	}
	res := make([][]int, 2)
	for i := 0; i < edgesLen; i++ {
		curUF := newUf(n)
		curVal := 0
		for j := 0; j < edgesLen; j++ {
			if j != i && curUF.union(edges[j][0], edges[j][1]) {
				curVal += edges[j][2]
			}
		}
		if curUF.size != 1 || curVal > minVal {
			res[0] = append(res[0], edges[i][3])
			continue
		}
		curUF = newUf(n)
		curUF.union(edges[i][0], edges[i][1])
		curVal = edges[i][2]
		for j := 0; j < edgesLen; j++ {
			if j != i && curUF.union(edges[j][0], edges[j][1]) {
				curVal += edges[j][2]
			}
		}
		if curVal == minVal {
			res[1] = append(res[1], edges[i][3])
		}
	}
	return res
}

func main() {
	n := 5
	edges := [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 2}, {0, 3, 2}, {0, 4, 3}, {3, 4, 3}, {1, 4, 6}}
	fmt.Println(findCriticalAndPseudoCriticalEdges(n, edges))
}
