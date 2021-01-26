package main

import (
	"fmt"
)

func exist(board [][]byte, word string) bool {
	row := len(board)
	col := len(board[0])
	visited := map[int]bool{}
	wordLen := len(word)
	dir4 := [][]int{{1, 0}, {0, 1}, {0, -1}, {-1, 0}}
	var calKey func(r, c int) int
	calKey = func(r, c int) int {
		return r*col + c
	}
	var dfs func(r, c, wordID int) bool
	dfs = func(r, c, wordID int) bool {
		if wordID == wordLen {
			return true
		}
		for _, m := range dir4 {
			newR, newC := r+m[0], c+m[1]
			newK := calKey(newR, newC)
			if 0 <= newR && newR < row &&
				0 <= newC && newC < col &&
				!visited[newK] && board[newR][newC] == word[wordID] {
				visited[newK] = true
				if dfs(newR, newC, wordID+1) {
					return true
				}
				visited[newK] = false
			}
		}
		return false
	}
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if board[r][c] == word[0] {
				visited = map[int]bool{}
				visited[calKey(r, c)] = true
				if dfs(r, c, 1) {
					return true
				}
			}
		}

	}
	return false
}

func main() {
	// board := [][]byte{
	// 	{'A', 'B', 'C', 'E'},
	// 	{'S', 'F', 'C', 'S'},
	// 	{'A', 'D', 'E', 'E'}}
	// word := "ABCCED"
	// board := [][]byte{{'a', 'a'}}
	// word := "aaa"
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'}}
	word := "ABCESEEEFS"
	fmt.Println(exist(board, word))
}
