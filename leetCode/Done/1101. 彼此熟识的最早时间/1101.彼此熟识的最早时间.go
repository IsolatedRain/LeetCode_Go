package main

import (
	"fmt"
	"sort"
)

func earliestAcq(logs [][]int, N int) int {
	// 按时间升序, 排序
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})

	// 父节点, 初始化为 自身
	p := make([]int, N, N)
	for i := range p {
		p[i] = i
	}

	var getRoot func(x int) int
	getRoot = func(x int) int {
		for p[x] != x {
			x = p[x]
		}
		return x
	}

	// 一共需要合并 N-1 次, 即两两认识一次.
	count := N - 1
	for _, v := range logs {
		xR := getRoot(v[1])
		yR := getRoot(v[2])
		if xR == yR {
			continue
		}
		p[xR] = yR
		count--
		if count == 0 {
			return v[0]
		}
	}

	return -1
}

func main() {
	logs := [][]int{{20190107, 2, 3}, {20190101, 0, 1}, {20190104, 3, 4}, {20190211, 1, 5}, {20190224, 2, 4}, {20190301, 0, 3}, {20190312, 1, 2}, {20190322, 4, 5}}
	n := 6
	fmt.Printf("输入: %v\nN: %v\n", logs, n)
	r := earliestAcq(logs, n)
	fmt.Printf("输出: %v\n", r)

}
