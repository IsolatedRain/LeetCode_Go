package main

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	n := len(intervals)
	res := 0
	R := intervals[0][1]
	for i := 1; i < n; i++ {
		if intervals[i][0] < R {
			res++
		} else {
			R = intervals[i][1]
		}
	}
	return res
}

func main() {
	intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	r := eraseOverlapIntervals(intervals)
	fmt.Println(r)
}
