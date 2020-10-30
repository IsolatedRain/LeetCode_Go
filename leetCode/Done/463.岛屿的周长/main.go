package main

import "fmt"

// DFS,每出界一次, 或者 遇到水, 边数 + 1
func islandPerimeter(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	res := 0
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var dfs func(x int, y int)
	dfs = func(x int, y int) {
		grid[x][y] = 2
		for _, m := range directions {
			nx := x + m[0]
			ny := y + m[1]
			if nx < 0 || ny < 0 || nx >= row || ny >= col || grid[nx][ny] == 0 {
				res++
				continue
			} else if grid[nx][ny] == 2 {
				continue
			} else {
				dfs(nx, ny)
			}
		}
	}
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if grid[r][c] == 1 {
				dfs(r, c)
			}
		}
	}
	return res
}

func main() {
	p := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
	r := islandPerimeter(p)
	fmt.Printf("%v\n", r)
}
