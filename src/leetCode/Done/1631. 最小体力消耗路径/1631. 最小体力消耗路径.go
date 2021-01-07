package main

import (
	"fmt"
)

// 二分
func minimumEffortPath(heights [][]int) int {
	row := len(heights)
	col := len(heights[0])
	if row == 1 && col == 1 {
		return 0
	}
	dir4 := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	visited := make([][]bool, row)
	for i := range visited {
		visited[i] = make([]bool, col)
	}
	var couldReach func(r, c, x int) bool
	couldReach = func(r, c, x int) bool {
		if r == row-1 && c == col-1 {
			return true
		}
		for _, m := range dir4 {
			nxR := r + m[0]
			nxC := c + m[1]
			if 0 <= nxR && nxR < row &&
				0 <= nxC && nxC < col &&
				!visited[nxR][nxC] &&
				abs(heights[r][c]-heights[nxR][nxC]) <= x {
				visited[nxR][nxC] = true
				if couldReach(nxR, nxC, x) {
					return true
				}
			}
		}
		return false
	}
	L := 1
	R := 1000000
	for L < R {
		mid := (L + R) >> 1
		for i := range visited {
			visited[i] = make([]bool, col)
		}
		if !couldReach(0, 0, mid) {
			L = mid + 1
		} else {
			R = mid
		}
	}
	return L
}

func abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	heights := [][]int{
		{1, 2, 3},
		{3, 8, 4},
		{5, 3, 5}}
	// heights := [][]int{{3}}
	// heights := [][]int{
	// 	{4, 3, 4, 10, 5, 5, 9, 2},
	// 	{10, 8, 2, 10, 9, 7, 5, 6},
	// 	{5, 8, 10, 10, 10, 7, 4, 2},
	// 	{5, 1, 3, 1, 1, 3, 1, 9},
	// 	{6, 4, 10, 6, 10, 9, 4, 6}}
	fmt.Println(minimumEffortPath(heights))
}

// 并查集
// func minimumEffortPath(heights [][]int) int {
// 	row := len(heights)
// 	col := len(heights[0])
// 	dir2 := [][]int{{1, 0}, {0, 1}}
// 	graph := [][]int{}
// 	for r := 0; r < row; r++ {
// 		for c := 0; c < col; c++ {
// 			curNode := r*col + c
// 			for _, m := range dir2 {
// 				nxR := r + m[0]
// 				nxC := c + m[1]
// 				if nxR < row && nxC < col {
// 					nxNode := nxR*col + nxC
// 					graph = append(graph, []int{curNode, nxNode, abs(heights[r][c] - heights[nxR][nxC])})
// 				}
// 			}
// 		}
// 	}
// 	sort.Slice(graph, func(i, j int) bool {
// 		return graph[i][2] < graph[j][2]
// 	})
// 	p := make([]int, row*col)
// 	for i := range p {
// 		p[i] = i
// 	}
// 	var getRoot func(x int) int
// 	getRoot = func(x int) int {
// 		if p[x] != p[p[x]] {
// 			p[x] = getRoot(p[x])
// 		}
// 		return p[x]
// 	}
// 	var union func(x, y int)
// 	union = func(x, y int) {
// 		xRoot := getRoot(x)
// 		yRoot := getRoot(y)
// 		if xRoot > yRoot {
// 			p[xRoot] = yRoot
// 		} else {
// 			p[yRoot] = xRoot
// 		}
// 	}

// 	end := row*col - 1
// 	for _, e := range graph {
// 		xRoot := getRoot(e[0])
// 		yRoot := getRoot(e[1])
// 		if xRoot != yRoot {
// 			union(xRoot, yRoot)
// 			if getRoot(0) == getRoot(end) {
// 				return e[2]
// 			}
// 		}
// 	}
// 	return 0
// }

// 最小生成树
// func minimumEffortPath(heights [][]int) int {
// 	row := len(heights)
// 	col := len(heights[0])
// 	mark := make([][]int, row)
// 	for i := range mark {
// 		mark[i] = make([]int, col)
// 		for j := range mark[i] {
// 			mark[i][j] = -1
// 		}
// 	}

// 	mark[0][0] = 0
// 	q := map[int]struct{}{0: {}}
// 	dir4 := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
// 	for len(q) > 0 {
// 		nxQ := map[int]struct{}{}
// 		for k := range q {
// 			curR := k / col
// 			curC := k % col
// 			for _, m := range dir4 {
// 				nxR := curR + m[0]
// 				nxC := curC + m[1]
// 				if 0 <= nxR && nxR < row && 0 <= nxC && nxC < col {
// 					v := max(abs(heights[curR][curC]-heights[nxR][nxC]), mark[curR][curC])
// 					if mark[nxR][nxC] == -1 || mark[nxR][nxC] > v {
// 						mark[nxR][nxC] = v
// 						nxQ[nxR*col+nxC] = struct{}{}
// 					}
// 				}
// 			}
// 		}
// 		q = nxQ
// 	}
// 	return mark[row-1][col-1]
// }
