package main

func rotate(matrix [][]int) {
	row := len(matrix)
	col := len(matrix[0])
	for r := 0; r < row; r++ {
		for c := r + 1; c < col; c++ {
			matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]
		}
	}
	for i := range matrix {
		reverse(matrix[i])
	}
}
func reverse(row []int) {
	n := len(row)
	L, R := 0, n-1
	for L < R {
		row[L], row[R] = row[R], row[L]
		L++
		R--
	}
}

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {4, 5, 6, 7}, {7, 8, 9, 0}, {1, 2, 3, 4}}
	rotate(matrix)
}
