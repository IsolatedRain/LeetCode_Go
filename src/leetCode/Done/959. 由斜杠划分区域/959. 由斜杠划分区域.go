package main

import "fmt"

func regionsBySlashes(grid []string) int {
	n := len(grid)
	p := make([]int, 4*n*n)
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
		xRoot, yRoot := getRoot(x), getRoot(y)
		if xRoot != yRoot {
			p[xRoot] = yRoot
		}
	}
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			curPos := 4 * (r*n + c)
			if grid[r][c] == '/' {
				union(curPos+0, curPos+3)
				union(curPos+1, curPos+2)
			}
			if grid[r][c] == '\\' {
				union(curPos+0, curPos+1)
				union(curPos+2, curPos+3)
			}
			if r+1 < n {
				union(curPos+2, curPos+4*n+0)
			}
			if r-1 >= 0 {
				union(curPos+0, curPos-4*n+2)
			}
			if c+1 < n {
				union(curPos+1, curPos+7)
			}
			if c-1 >= 0 {
				union(curPos+3, curPos-3)
			}
		}
	}
	res := 0
	for i, v := range p {
		if i == v {
			res++
		}
	}
	return res
}

func main() {
	grid := []string{
		"\\/",
		"/\\"}
	fmt.Println(regionsBySlashes(grid))
}
