package main

import (
	"fmt"
)

func pathsWithMaxScore(board []string) []int {
	row, col := len(board), len(board[0])
	b := make([][]byte, len(board))
	for i := range board {
		b[i] = []byte(board[i])
	}
	path := make([][]int, row+1)
	score := make([][]int, row+1)
	for i := range path {
		path[i] = make([]int, col+1)
		score[i] = make([]int, col+1)
	}
	path[row-1][col-1] = 1
	mod := 1000000007
	for r := row - 1; r >= 0; r-- {
		for c := col - 1; c >= 0; c-- {
			if b[r][c] != 'X' {
				curMaxScore := max(score[r+1][c], score[r+1][c+1], score[r][c+1])
				curPaths := 0
				score[r][c] = curMaxScore
				if isnum(b[r][c]) {
					score[r][c] += int(b[r][c] - '0')
				}
				if score[r+1][c] == curMaxScore {
					curPaths += path[r+1][c]
				}
				if score[r+1][c+1] == curMaxScore {
					curPaths += path[r+1][c+1]
				}
				if score[r][c+1] == curMaxScore {
					curPaths += path[r][c+1]
				}
				path[r][c] += curPaths % mod
				score[r][c] %= mod
			}
		}
	}
	for _, s := range score {
		fmt.Println(s)
	}
	fmt.Println("---")
	for _, p := range path {
		fmt.Println(p)
	}
	if path[0][0] == 0 {
		return []int{0, 0}
	}
	return []int{score[0][0], path[0][0]}
}
func isnum(x byte) bool {
	if '0' <= x && x <= '9' {
		return true
	}
	return false
}
func max(a ...int) int {
	r := a[0]
	for _, v := range a[1:] {
		if v > r {
			r = v
		}
	}
	return r
}

func main() {
	// board := []string{
	// 	"E23",
	// 	"2X2",
	// 	"12S"}
	board := []string{
		"E2XX6",
		"58299",
		"74213",
		"X6692",
		"7681S"}
	fmt.Println(pathsWithMaxScore(board))
}
