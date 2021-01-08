package main

import "fmt"

func maxDistance(grid [][]int) int {
	dir4 := [][]int{{0, 1}, {-1, 0}, {1, 0}, {0, -1}}
	row := len(grid)
	col := len(grid[0])
	visited := make([][]bool, row)
	for i := range visited {
		visited[i] = make([]bool, col)
	}
	q := [][]int{}
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if grid[r][c] == 1 {
				q = append(q, []int{r, c, 0})
				visited[r][c] = true
			}
		}
	}
	if len(q) == 0 || len(q) == row*col {
		return -1
	}
	res := -1
	for len(q) > 0 {
		curLen := len(q)
		for i := 0; i < curLen; i++ {
			r, c, steps := q[i][0], q[i][1], q[i][2]
			res = max(res, steps)
			for _, m := range dir4 {
				newR, newC := r+m[0], c+m[1]
				if 0 <= newR && newR < row &&
					0 <= newC && newC < col &&
					!visited[newR][newC] {
					visited[newR][newC] = true
					q = append(q, []int{newR, newC, steps + 1})
				}
			}
		}
		q = q[curLen:]
	}
	return res
}
func max(x ...int) int {
	r := x[0]
	for _, v := range x[1:] {
		if v > r {
			r = v
		}
	}
	return r
}

func main() {
	// grid := [][]int{
	// 	{1, 0, 1},
	// 	{0, 0, 0},
	// 	{0, 0, 0}}
	// grid := [][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}
	// grid := [][]int{{1, 1}, {1, 1}}
	grid := [][]int{
		{1, 0, 0, 0, 0, 1, 0, 0, 0, 1},
		{1, 1, 0, 1, 1, 1, 0, 1, 1, 0},
		{0, 1, 1, 0, 1, 0, 0, 1, 0, 0},
		{1, 0, 1, 0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 1, 1, 0, 1, 1},
		{0, 0, 1, 0, 0, 1, 0, 1, 0, 1},
		{0, 0, 0, 1, 1, 1, 1, 0, 0, 1},
		{0, 1, 0, 0, 1, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 1, 1, 1, 0, 0},
		{1, 1, 0, 1, 1, 1, 1, 1, 0, 0}}
	fmt.Println(maxDistance(grid))
}
