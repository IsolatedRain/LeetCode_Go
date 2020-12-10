package main

import "fmt"

func minPathSum(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	last := make([]int, col)
	cur := make([]int, col)
	last[0] = grid[0][0]
	for i := 1; i < col; i++ {
		last[i] += grid[0][i] + last[i-1]
	}
	for r := 1; r < row; r++ {
		cur[0] = grid[r][0] + last[0]
		for c := 1; c < col; c++ {
			cur[c] = min(cur[c-1], last[c]) + grid[r][c]
		}
		last = cur
		cur = last
	}
	return last[col-1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	// grid := [][]int{{1, 2}, {1, 1}}
	r := minPathSum(grid)
	fmt.Println(r)
}
