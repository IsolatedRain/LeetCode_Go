package main

import (
	"fmt"
)

func shortestDistanceColor(colors []int, queries [][]int) []int {
	n := len(colors)
	inf := n + 1
	left, right := make([][4]int, n), make([][4]int, n)
	res := []int{}
	for i := range left {
		left[i] = [4]int{0, inf, inf, inf}
		right[i] = [4]int{0, inf, inf, inf}
	}
	left[0][colors[0]] = 0
	for i := 1; i < n; i++ {
		if left[i-1][1] != inf {
			left[i][1] = left[i-1][1] + 1
		}
		if left[i-1][2] != inf {
			left[i][2] = left[i-1][2] + 1
		}
		if left[i-1][3] != inf {
			left[i][3] = left[i-1][3] + 1
		}
		left[i][colors[i]] = 0
	}
	right[n-1][colors[n-1]] = 0
	for i := n - 2; i >= 0; i-- {
		if right[i+1][1] != inf {
			right[i][1] = right[i+1][1] + 1
		}
		if right[i+1][2] != inf {
			right[i][2] = right[i+1][2] + 1
		}
		if right[i+1][3] != inf {
			right[i][3] = right[i+1][3] + 1
		}
		right[i][colors[i]] = 0
	}
	for _, q := range queries {
		L := left[q[0]][q[1]]
		R := right[q[0]][q[1]]
		res = append(res, min(L, R))
	}
	return res
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	fmt.Println(shortestDistanceColor([]int{1, 1, 2, 1, 3, 2, 2, 3, 3}, [][]int{{1, 3}, {2, 2}, {6, 1}}))
}
