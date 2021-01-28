package main

import (
	"fmt"
	"sort"
)

func pondSizes(land [][]int) []int {
	dir8 := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	row, col := len(land), len(land[0])
	res := []int{}
	mark := make([][]bool, row)
	for i := range mark {
		mark[i] = make([]bool, col)
	}
	var dfs func(r, c int)
	var curSize int
	dfs = func(r, c int) {
		for _, m := range dir8 {
			nr, nc := r+m[0], c+m[1]
			if 0 <= nr && nr < row && 0 <= nc && nc < col && land[nr][nc] == 0 && !mark[nr][nc] {
				mark[nr][nc] = true
				dfs(nr, nc)
				curSize++
			}
		}
	}
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if land[r][c] == 0 && !mark[r][c] {
				curSize = 1
				mark[r][c] = true
				dfs(r, c)
				res = append(res, curSize)
			}
		}
	}
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	return res
}

func main() {
	land := [][]int{
		{0, 2, 1, 0},
		{0, 1, 0, 1},
		{1, 1, 0, 1},
		{0, 1, 0, 1}}
	fmt.Println(pondSizes(land))
}
