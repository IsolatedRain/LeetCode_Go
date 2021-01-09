package main

import (
	"fmt"
	"sort"
)

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] > intervals[j][1])
	})
	toDeleteNum := 0
	R := -1
	for _, i := range intervals {
		if i[1] <= R {
			toDeleteNum++
		} else {
			R = i[1]
		}
	}
	return len(intervals) - toDeleteNum
}

func main() {
	intervals := [][]int{{1, 4}, {3, 6}, {2, 8}}
	fmt.Println(removeCoveredIntervals(intervals))
}
