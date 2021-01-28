package main

import (
	"fmt"
)

func getFood(grid [][]byte) int {
	row, col := len(grid), len(grid[0])
	var start []int
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if grid[r][c] == '*' {
				start = []int{r, c, 0}
				break
			}
		}
	}
	q := [][]int{start}
	visited := map[int]bool{}
	dir4 := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(q) > 0 {
		curLen := len(q)
		for _, v := range q {
			r, c, steps := v[0], v[1], v[2]
			if grid[r][c] == '#' {
				return steps
			}
			for _, m := range dir4 {
				nr, nc := r+m[0], c+m[1]
				if 0 <= nr && nr < row && 0 <= nc && nc < col &&
					!visited[nr*col+nc] && grid[nr][nc] != 'X' {
					visited[nr*col+nc] = true
					q = append(q, []int{nr, nc, steps + 1})
				}
			}
		}
		q = q[curLen:]
	}
	return -1
}

func main() {
	fmt.Println(getFood([][]byte{
		{'X', 'X', 'X', 'X', 'X', 'X'},
		{'X', '*', 'O', 'O', 'O', 'X'},
		{'X', 'O', 'O', '#', 'O', 'X'},
		{'X', 'X', 'X', 'X', 'X', 'X'}}))
}
