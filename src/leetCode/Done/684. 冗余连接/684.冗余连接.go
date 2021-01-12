package main

import "fmt"

func findRedundantConnection(edges [][]int) []int {
	p := make([]int, len(edges)+1)
	for i := range p {
		p[i] = i
	}
	var getRoot func(x int) int
	getRoot = func(x int) int {
		// 寻找父节点
		if p[x] != p[p[x]] {
			p[x] = getRoot(p[x])
		}
		return p[x]
	}
	// 合并
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
	for _, e := range edges {
		e0 := getRoot(e[0])
		e1 := getRoot(e[1])
		if e0 != e1 {
			union(e0, e1)
		} else {
			return e
		}
	}
	return []int{}
}

func main() {
	// p := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}
	// p := [][]int{{1, 2}, {1, 3}, {2, 3}}
	p := [][]int{{1, 5}, {3, 4}, {3, 5}, {4, 5}, {2, 4}}
	fmt.Printf("输入: %v\n", p)
	r := findRedundantConnection(p)
	fmt.Printf("输出: %v\n", r)

}
