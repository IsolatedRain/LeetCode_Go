package main

import "fmt"

func setZeroes(matrix [][]int) [][]int {
	row, col := len(matrix), len(matrix[0])
	firRow := false
	firCol := false
	for i := 0; i < row; i++ {
		if matrix[i][0] == 0 {
			firCol = true
			break
		}
	}
	for i := 0; i < col; i++ {
		if matrix[0][i] == 0 {
			firRow = true
			break
		}
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < row; i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < col; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for i := 1; i < col; i++ {
		if matrix[0][i] == 0 {
			for j := 0; j < row; j++ {
				matrix[j][i] = 0
			}
		}
	}
	if firCol {
		for i := 0; i < row; i++ {
			matrix[i][0] = 0
		}
	}
	if firRow {
		for i := 0; i < col; i++ {
			matrix[0][i] = 0
		}
	}
	return matrix
}

func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	fmt.Println(setZeroes(matrix))
}
