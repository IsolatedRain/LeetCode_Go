package main

import (
	"fmt"
	"sort"
)

func minimumCost(N int, connections [][]int) int {
	// 初始化 根节点为自身
	p := make([]int, N+1)
	for i := range p {
		p[i] = i
	}
	// 按 花费升序, 排序
	sort.Slice(connections, func(i, j int) bool {
		return connections[i][2] < connections[j][2]
	})

	var getRoot func(x int) int
	getRoot = func(x int) int {
		for p[x] != x {
			x = p[x]
		}
		return x
	}
	// 联通N-1次
	count := N - 1
	curCost := 0
	for _, v := range connections {
		xRoot := getRoot(v[0])
		yRoot := getRoot(v[1])
		if xRoot == yRoot {
			continue
		}
		p[xRoot] = yRoot
		curCost += v[2]
		count--
		if count == 0 {
			return curCost
		}
	}

	return -1
}

func main() {
	N := 3
	conections := [][]int{{1, 2, 5}, {1, 3, 6}, {2, 3, 1}}
	fmt.Printf("输入: %v\n", N)
	fmt.Printf("     %v\n", conections)
	r := minimumCost(N, conections)
	fmt.Printf("输出: %v\n", r)
}
