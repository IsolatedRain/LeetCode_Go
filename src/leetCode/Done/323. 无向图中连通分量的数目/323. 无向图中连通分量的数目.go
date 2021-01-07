package main

import "fmt"

// DSU 并查集
type DSU struct {
	p []int
}

func (d *DSU) getRoot(x int) int {
	if d.p[x] != d.p[d.p[x]] {
		d.p[x] = d.getRoot(d.p[x])
	}
	return d.p[x]
}
func (d *DSU) union(x, y int) {
	xRoot := d.getRoot(x)
	yRoot := d.getRoot(y)
	if xRoot > yRoot {
		d.p[xRoot] = yRoot
	} else {
		d.p[yRoot] = xRoot
	}
}

// NewDSU 并查集
func NewDSU(n int) *DSU {
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}
	return &DSU{p}
}

func countComponents(n int, edges [][]int) int {
	p := NewDSU(n)
	for _, e := range edges {
		p.union(e[0], e[1])
	}
	res := 0
	for i, v := range p.p {
		if i == v {
			res++
		}
	}
	return res
}

func main() {
	n := 5
	edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}
	fmt.Println(countComponents(n, edges))
}
