package main

import "fmt"

func numIslands(grid [][]byte) int {
	row := len(grid)
	col := len(grid[0])
	p := make([]int, row*col)
	for i := range p {
		p[i] = i
	}
	var getRoot func(x int) int
	getRoot = func(x int) int {
		if p[x] != p[p[x]] {
			p[x] = getRoot(p[x])
		}
		return p[x]
	}
	var union func(x, y int)
	union = func(x, y int) {
		xRoot := getRoot(x)
		yRoot := getRoot(y)
		if xRoot > yRoot {
			p[xRoot] = yRoot
		} else {
			p[yRoot] = xRoot
		}
	}
	dir2 := [][]int{{0, 1}, {1, 0}}
	visited := make([]bool, row*col)
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if grid[r][c] == '0' {
				p[r*col+c] = -1
				continue
			}
			if grid[r][c] == '1' && !visited[r*col+c] {
				visited[r*col+c] = true
				for _, m := range dir2 {
					newR := r + m[0]
					newC := c + m[1]
					if newR < row &&
						newC < col &&
						grid[newR][newC] == '1' &&
						!visited[newR*col+newC] {
						union(r*col+c, newR*col+newC)
					}
				}
			}
		}
	}
	res := 0
	for i := range p {
		if p[i] == i {
			res++
		}
	}
	return res
}

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'}}
	fmt.Println(numIslands(grid))
}
