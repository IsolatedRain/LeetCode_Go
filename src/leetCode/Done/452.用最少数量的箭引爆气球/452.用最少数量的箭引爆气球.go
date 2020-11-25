package main

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	curRight := points[0][1]
	cnt := 1
	for _, p := range points[1:] {
		if p[0] > curRight {
			cnt++
			curRight = p[1]
		}
	}
	return cnt
}

func main() {
	points := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	r := findMinArrowShots(points)
	fmt.Printf("%v\n", r)
}
