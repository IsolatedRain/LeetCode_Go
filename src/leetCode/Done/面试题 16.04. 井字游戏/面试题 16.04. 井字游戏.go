package main

import "fmt"

func tictactoe(board []string) string {
	n := len(board)

	var check func([][]int) string
	check = func(arr [][]int) string {
		countX := 0
		countO := 0
		for i := range arr {
			if board[arr[i][0]][arr[i][1]] == 'X' {
				countX++
			} else if board[arr[i][0]][arr[i][1]] == 'O' {
				countO++
			}
		}
		if countX == n {
			return "X"
		}
		if countO == n {
			return "O"
		}
		return "Draw"
	}

	dia1 := [][]int{}
	dia2 := [][]int{}
	for i := 0; i < n; i++ {
		dia1 = append(dia1, []int{i, i})
		dia2 = append(dia2, []int{i, n - i - 1})
	}
	rDia1 := check(dia1)
	rDia2 := check(dia2)
	if rDia1 == "X" || rDia2 == "X" {
		return "X"
	}
	if rDia1 == "O" || rDia2 == "O" {
		return "O"
	}

	for i := 0; i < n; i++ {
		curCol := [][]int{}
		curRow := [][]int{}
		for j := 0; j < n; j++ {
			curRow = append(curRow, []int{i, j})
			curCol = append(curCol, []int{j, i})
		}
		rRow := check(curRow)
		rCol := check(curCol)
		if rRow == "X" || rCol == "X" {
			return "X"
		}
		if rRow == "O" || rCol == "O" {
			return "O"
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == ' ' {
				return "Pending"
			}
		}
	}
	return "Draw"
}

func main() {
	board := []string{
		"OOX",
		"XXO",
		"OX "}
	r := tictactoe(board)
	fmt.Println(r)
}
