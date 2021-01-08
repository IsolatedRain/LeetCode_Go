package main

import "fmt"

func makeConnected(n int, connections [][]int) int {
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
		xRoot := getRoot(x)
		yRoot := getRoot(y)
		if xRoot > yRoot {
			p[xRoot] = yRoot
		} else {
			p[yRoot] = xRoot
		}
	}
	for _, c := range connections {
		union(c[0], c[1])
	}
	lines := 0
	for i := range p {
		if p[i] == i {
			lines++
		}
	}
	if len(connections) >= n-1 {
		return lines - 1
	}
	return -1
}

func main() {
	n := 6
	connections := [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}}
	fmt.Println(makeConnected(n, connections))
}
