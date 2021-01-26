package main

func solve(board [][]byte) {
	row, col := len(board), len(board[0])
	dir4 := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	var dfs func(r, c int)
	dfs = func(r, c int) {
		board[r][c] = '#'
		for _, m := range dir4 {
			newR, newC := r+m[0], c+m[1]
			if 0 <= newR && newR < row &&
				0 <= newC && newC < col && board[newR][newC] == 'O' {
				dfs(newR, newC)
			}
		}
	}

	for r := 0; r < row; r++ {
		if board[r][0] == 'O' {
			dfs(r, 0)
		}
		if board[r][col-1] == 'O' {
			dfs(r, col-1)
		}
	}
	for c := 0; c < col; c++ {
		if board[0][c] == 'O' {
			dfs(0, c)
		}
		if board[row-1][c] == 'O' {
			dfs(row-1, c)
		}
	}
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if board[r][c] == 'O' {
				board[r][c] = 'X'
			} else if board[r][c] == '#' {
				board[r][c] = 'O'
			}
		}
	}
	return
}
func main() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'}}
	solve(board)
}
