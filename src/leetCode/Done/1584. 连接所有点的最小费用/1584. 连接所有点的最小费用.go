package main

import (
	"fmt"
	"sort"
)

func minCostConnectPoints(points [][]int) int {
	n := len(points)

	var abs func(x int) int
	abs = func(x int) int {
		if x < 0 {
			x *= -1
		}
		return x
	}

	var calDist func(x, y []int) int
	calDist = func(x, y []int) int {
		return abs(x[0]-y[0]) + abs(x[1]-y[1])
	}

	graph := [][]int{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			graph = append(graph, []int{i, j, calDist(points[i], points[j])})
		}
	}
	sort.Slice(graph, func(i, j int) bool { return graph[i][2] < graph[j][2] })

	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}

	var getRoot func(x int) int
	getRoot = func(x int) int {
		for p[x] != x {
			p[x] = p[p[x]]
		}
		return x
	}

	res := 0
	count := 0
	for _, v := range graph {
		xRoot := getRoot(v[0])
		yRoot := getRoot(v[1])
		if xRoot != yRoot {
			res += v[2]
			count++
			if count == n-1 {
				break
			}
			p[xRoot] = p[yRoot]
		}
	}
	return res
}

func main() {
	points := [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}
	fmt.Println(minCostConnectPoints(points))
}
