package main

import (
	"fmt"
	"sort"
)

func swimInWater(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	dir4 := [][]int{{0, 1}, {1, 0}}
	graph := [][]int{}
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			for _, m := range dir4 {
				nr, nc := r+m[0], c+m[1]
				if 0 <= nr && nr < row && 0 <= nc && nc < col {
					graph = append(graph, []int{r*col + c, nr*col + nc, max(grid[r][c], grid[nr][nc])})
				}
			}
		}
	}
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
		xRoot, yRoot := getRoot(x), getRoot(y)
		if xRoot != yRoot {
			p[xRoot] = yRoot
		}
	}
	sort.Slice(graph, func(i, j int) bool { return graph[i][2] < graph[j][2] })
	start, end := 0, row*col-1
	for _, e := range graph {
		xRoot, yRoot := getRoot(e[0]), getRoot(e[1])
		if xRoot != yRoot {
			union(e[0], e[1])
		}
		if getRoot(start) == getRoot(end) {
			return e[2]
		}
	}
	return -1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 二分
// func swimInWater(grid [][]int) int {
// 	row, col, visited := len(grid), len(grid[0]), map[int]bool{}
// 	visited[0] = true
// 	dir4 := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
// 	q := [][]int{{0, 0, grid[0][0]}}
// 	for len(q) > 0 {
// 		v := q[0]
// 		q = q[1:]
// 		r, c, depth := v[0], v[1], v[2]
// 		if r == row-1 && c == col-1 {
// 			return depth
// 		}
// 		for _, m := range dir4 {
// 			nr, nc := r+m[0], c+m[1]
// 			if 0 <= nr && nr < row && 0 <= nc && nc < col && !visited[nr*col+nc] {
// 				visited[nr*col+nc] = true
// 				curDepth := max(depth, grid[nr][nc])
// 				i := sort.Search(len(q), func(i int) bool { return q[i][2] >= curDepth })
// 				q = append(q[:i], append([][]int{{nr, nc, curDepth}}, q[i:]...)...)
// 			}
// 		}
// 	}
// 	return -1
// }

func main() {
	fmt.Println(swimInWater([][]int{{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}, {10, 9, 8, 7, 6}}))
}
