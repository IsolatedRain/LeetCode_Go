package main

import "fmt"

func gameOfLife(board [][]int) [][]int {
	dir8 := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	row, col := len(board), len(board[0])
	var count func(x, y int) int
	count = func(x, y int) int {
		cnt := 0
		for _, m := range dir8 {
			newX := x + m[0]
			newY := y + m[1]
			if 0 <= newX && newX < row &&
				0 <= newY && newY < col && board[newX][newY] == 1 {
				cnt++
			}
		}
		return cnt
	}

	status := make([][]int, row)
	for i := range status {
		status[i] = append(status[i], board[i]...)
	}

	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			liveNums := count(r, c)
			if board[r][c] == 1 {
				if 2 <= liveNums && liveNums <= 3 {
					status[r][c] = 1
				} else {
					status[r][c] = 0
				}
			} else {

				if liveNums == 3 {
					fmt.Println(r, c, liveNums)
					status[r][c] = 1
				}
			}
		}
	}
	return status
}
func main() {
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0}}
	fmt.Println(gameOfLife(board))
}
