// https://leetcode-cn.com/problems/spiral-matrix-ii/
package main

import "fmt"

func generateMatrix(n int) [][]int {
	t, b, l, r := 0, n, 0, n
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	curNum, maxNum := 1, n*n+1
	for curNum < maxNum {
		for i := l; i < r; i++ {
			matrix[t][i] = curNum
			curNum++
		}
		t++
		for i := t; i < b; i++ {
			matrix[i][r-1] = curNum
			curNum++
		}
		r--
		for i := r-1; i >= l; i-- {
			matrix[b-1][i] = curNum
			curNum++
		}
		b--
		for i := b-1; i >= t; i-- {
			matrix[i][l] = curNum
			curNum++
		}
		l++
	}
	return matrix
}

func main() {
	fmt.Println(generateMatrix(5))
}
