package main

import "fmt"

func findRedundantConnection(edges [][]int) []int {
	p := make(map[int]int)

	var getRoot func(x int) int
	getRoot = func(x int) int {
		// 第一次遇到, 初始化 为自身
		if _, ok := p[x]; !ok {
			p[x] = x
		}
		// 寻找父节点
		for p[x] != x {
			x = p[x]
		}
		return x
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
			return []int{e[0], e[1]}
		}

	}
	return []int{}
}
func main() {
	p := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}
	fmt.Printf("输入: %v\n", p)
	r := findRedundantConnection(p)
	fmt.Printf("输出: %v\n", r)

}
