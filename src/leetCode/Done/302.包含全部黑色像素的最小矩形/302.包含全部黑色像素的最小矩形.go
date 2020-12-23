package main

import "fmt"

func minArea(image [][]byte, x int, y int) int {
	row := len(image)
	col := len(image[0])
	dir4 := [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
	visited := map[int]bool{x*col + y: true}
	var dfs func(r, c int)
	L, R, T, D := col+1, -1, row+1, -1
	dfs = func(r, c int) {
		L = min(L, c)
		R = max(R, c)
		T = min(T, r)
		D = max(D, r)
		for _, m := range dir4 {
			nxR := m[0] + r
			nxC := m[1] + c
			curK := nxR*col + nxC
			if 0 <= nxR && nxR < row && 0 <= nxC && nxC < col && image[nxR][nxC] == '1' && !visited[curK] {
				visited[curK] = true
				dfs(nxR, nxC)
			}
		}
	}
	dfs(x, y)
	return (R - L + 1) * (D - T + 1)
}
func min(x, y int) int {
	if x > y {
		return y
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
	image := [][]byte{
		{'0', '0', '1', '0'},
		{'0', '1', '1', '0'},
		{'0', '1', '0', '0'}}
	x := 0
	y := 2
	r := minArea(image, x, y)
	fmt.Println(r)
}
