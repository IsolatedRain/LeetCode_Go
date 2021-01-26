package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	res := []int{}
	for len(matrix) > 0 {
		nextMat := [][]int{}
		res = append(res, matrix[0]...)
		matrix = matrix[1:]
		if len(matrix) == 0 {
			break
		}
		for c := len(matrix[0]) - 1; c >= 0; c-- {
			curCol := make([]int, len(matrix))
			for r := 0; r < len(matrix); r++ {
				curCol[r] = matrix[r][c]
			}
			nextMat = append(nextMat, curCol)
		}
		matrix = nextMat
	}
	return res
}

func reverse(mat [][]int) [][]int {
	L, R := len(mat), len(mat)-1
	for L < R {
		mat[L], mat[R] = mat[R], mat[L]
		L++
		R--
	}
	return mat
}
func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}
	fmt.Println(spiralOrder(matrix))
}
