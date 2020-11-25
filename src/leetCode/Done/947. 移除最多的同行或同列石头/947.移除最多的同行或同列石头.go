package main

import "fmt"

func removeStones(stones [][]int) int {
	p := make(map[int]int)
	var getRoot func(x int) int
	getRoot = func(x int) int {
		if _, ok := p[x]; !ok {
			p[x] = x
		}
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
	for _, pos := range stones {
		xR := getRoot(pos[0])
		yR := getRoot(pos[1] + 10000)
		union(xR, yR)
	}
	root := make(map[int]struct{}, 0)
	for _, pos := range stones {
		k := getRoot(pos[0])
		root[k] = struct{}{}
	}
	return len(stones) - len(root)
}

func main() {
	stones := [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}
	fmt.Printf("输入: %v\n", stones)
	r := removeStones(stones)
	fmt.Printf("输出: %v\n", r)

}
