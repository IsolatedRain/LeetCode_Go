package main

import "fmt"

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	rows := make([]int, n)
	cols := make([]int, n)
	for i := range grid {
		rows[i] = max(grid[i]...)
		curCol := []int{}
		for j := 0; j < n; j++ {
			curCol = append(curCol, grid[j][i])
		}
		cols[i] = max(curCol...)
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res += min(rows[i], cols[j]) - grid[i][j]
		}
	}
	return res
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x ...int) int {
	r := x[0]
	for _, num := range x {
		if r < num {
			r = num
		}
	}
	return r
}

func main() {
	grid := [][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}
	r := maxIncreaseKeepingSkyline(grid)
	fmt.Println(r)
}
