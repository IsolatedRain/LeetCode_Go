package main

import "fmt"

func pacificAtlantic(matrix [][]int) [][]int {
	row, col := len(matrix), len(matrix[0])
	markP := make([][]byte, row)
	markA := make([][]byte, row)
	for i := range markP {
		markP[i] = make([]byte, col)
		markA[i] = make([]byte, col)
	}
	for i := 0; i < col; i++ {
		markP[0][i] = 'P'
		markA[row-1][i] = 'A'
	}
	for i := 0; i < row; i++ {
		markP[i][0] = 'P'
		markA[i][col-1] = 'A'
	}
	dir4 := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	var dfs func(r, c int, mk byte, mxt [][]byte)
	dfs = func(r, c int, mk byte, mtx [][]byte) {
		mtx[r][c] = mk
		for _, m := range dir4 {
			newR, newC := r+m[0], c+m[1]
			if 0 <= newR && newR < row &&
				0 <= newC && newC < col && mtx[newR][newC] != mk && matrix[r][c] <= matrix[newR][newC] {
				dfs(newR, newC, mk, mtx)
			}
		}
	}
	res := [][]int{}
	for c := 0; c < col; c++ {
		dfs(0, c, 'P', markP)
		dfs(row-1, c, 'A', markA)
	}
	for r := 0; r < row; r++ {
		dfs(r, 0, 'P', markP)
		dfs(r, col-1, 'A', markA)
	}
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if markP[r][c] == 'P' && markA[r][c] == 'A' {
				res = append(res, []int{r, c})
			}
		}
	}
	return res
}

func main() {
	matrix := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4}}
	fmt.Println(pacificAtlantic(matrix))
}
