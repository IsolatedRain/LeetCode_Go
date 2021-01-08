package main

import "fmt"

func countServers(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	r := make([]bool, row)
	c := make([]bool, col)
	for i := 0; i < row; i++ {
		count := 0
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				count++
			}
		}
		if count > 1 {
			r[i] = true
		}
	}
	for i := 0; i < col; i++ {
		count := 0
		for j := 0; j < row; j++ {
			if grid[j][i] == 1 {
				count++
			}
		}
		if count > 1 {
			c[i] = true
		}
	}
	res := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				if r[i] || c[j] {
					res++
				}
			}
		}
	}
	return res
}

func main() {
	grid := [][]int{{1, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}
	fmt.Println(countServers(grid))
}
