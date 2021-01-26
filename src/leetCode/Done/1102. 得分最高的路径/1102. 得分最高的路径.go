package main

import (
	"fmt"
	"sort"
)

func maximumMinimumPath(A [][]int) int {
	row, col := len(A), len(A[0])
	points := make([][]int, row*col)
	mark := make([]bool, row*col)
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			i := r*col + c
			points[i] = []int{A[r][c], r, c}
		}
	}
	sort.Slice(points, func(i, j int) bool { return points[i][0] > points[j][0] })

	root := make([]int, row*col)
	for i := range root {
		root[i] = i
	}

	var getRoot func(x int) int
	getRoot = func(x int) int {
		if root[x] != root[root[x]] {
			root[x] = getRoot(root[x])
		}
		return root[x]
	}
	var union func(x, y int)
	union = func(x, y int) {
		xRoot, yRoot := getRoot(x), getRoot(y)
		if xRoot != yRoot {
			root[xRoot] = yRoot
		}
	}
	dir4 := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	res := min(A[0][0], A[row-1][col-1])
	start, end := 0, row*col-1
	for _, v := range points {
		r, c := v[1], v[2]
		i := r*col + c
		res = min(res, v[0])
		mark[i] = true
		for _, m := range dir4 {
			nR, nC := r+m[0], c+m[1]
			if 0 <= nR && nR < row && 0 <= nC && nC < col && mark[nR*col+nC] {
				union(i, nR*col+nC)
				if getRoot(start) == getRoot(end) {
					return res
				}
			}
		}
	}
	return res
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	fmt.Println(maximumMinimumPath([][]int{
		{3, 4, 6, 3, 4},
		{0, 2, 1, 1, 7},
		{8, 8, 3, 2, 7},
		{3, 2, 4, 9, 8},
		{4, 1, 2, 0, 0},
		{4, 6, 5, 4, 3}}))
}
