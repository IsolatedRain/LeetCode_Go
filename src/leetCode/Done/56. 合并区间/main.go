package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}

	// 排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	curL, curR := intervals[0][0], intervals[0][1]
	n := len(intervals)
	res := [][]int{{curL, curR}}
	for i := 1; i < n; i++ {
		// 和前段 不相连
		if intervals[i][0] > curR {
			res = append(res, intervals[i])
			curL, curR = intervals[i][0], intervals[i][1]
		} else if curR < intervals[i][1] { // 和前段 交叉
			res = append(res[:len(res)-1], []int{curL, intervals[i][1]})
			curR = intervals[i][1]
		}
	}
	return res
}

func main() {
	// p := [][]int{{1, 3}, {8, 10}, {2, 6}, {15, 18}}
	p := [][]int{{1, 4}, {4, 5}}
	fmt.Printf("%v\n", p)
	r := merge(p)
	fmt.Printf("输入:%v\n输出: %v \n", p, r)
}
