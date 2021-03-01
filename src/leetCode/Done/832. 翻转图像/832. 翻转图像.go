package main

import "fmt"

func flipAndInvertImage(A [][]int) [][]int {
	row, col := len(A), len(A[0])
	for i := range A {
		L, R := 0, col-1
		for L < R {
			A[i][L], A[i][R] = A[i][R], A[i][L]
			L++
			R--
		}
	}
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			A[r][c] ^= 1
		}
	}
	return A
}
func main() {
	A := [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}
	fmt.Println(flipAndInvertImage(A))
}
