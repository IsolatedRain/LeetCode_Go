package main

import "fmt"

func uniquePathsIII(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	visited := make([][]int, row, row)
	startR, startC, tarR, tarC, tar := -1, -1, -1, -1, row*col
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if grid[r][c] == 1 {
				startR, startC = r, c
			} else if grid[r][c] == 2 {
				tarR, tarC = r, c
			} else if grid[r][c] == -1 {
				tar--
			}
		}
	}
	fmt.Println(startR, startC, tarR, tarC, tar)
	for i := range visited {
		visited[i] = make([]int, col)
	}
	visited[startR][startC] = 1
	for _, v := range grid {
		fmt.Println(v)
	}
	dir4 := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	res := 0
	var dfs func(r, c, step int)
	dfs = func(r, c, step int) {
		if r == tarR && c == tarC {
			if step == tar {
				res++
			}
			return
		}
		for _, m := range dir4 {
			nxR := r + m[0]
			nxC := c + m[1]
			if 0 <= nxR && nxR < row && 0 <= nxC && nxC < col && grid[nxR][nxC] != -1 && visited[nxR][nxC] == 0 {
				visited[nxR][nxC] = 1
				dfs(nxR, nxC, step+1)
				visited[nxR][nxC] = 0
			}
		}
		return
	}
	dfs(startR, startC, 1)
	return res
}

func main() {
	grid := [][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 2}}
	// grid := [][]int{{1, -1}, {0, 2}}
	r := uniquePathsIII(grid)
	fmt.Println(r)
}
