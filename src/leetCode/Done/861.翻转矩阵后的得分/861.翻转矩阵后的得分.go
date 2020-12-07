package main

import "fmt"

// 不翻转矩阵，直接计算每列的贡献。
// 1<<(c-1)
func matrixScore(A [][]int) int {
	row := len(A)
	col := len(A[0])
	res := row * (1 << (col - 1))
	for c := 1; c < col; c++ {
		cnt1 := 0
		for r := 0; r < row; r++ {
			if A[r][c] == A[r][0] {
				cnt1++
			}
		}
		if cnt1 < row-cnt1 {
			cnt1 = row - cnt1
		}
		res += cnt1 * (1 << (col - c - 1))
	}
	return res
}

func main() {
	A := [][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}
	// A := [][]int{{0, 1}, {1, 1}}
	r := matrixScore(A)
	fmt.Println(r)
}

// 翻转了原矩阵
// func matrixScore(A [][]int) int {
// 	row := len(A)
// 	col := len(A[0])
// 	var toggleRow func(r int)
// 	toggleRow = func(r int) {
// 		for c, num := range A[r] {
// 			A[r][c] = num ^ 1
// 		}
// 	}
// 	var toggleCol func(c int)
// 	toggleCol = func(c int) {
// 		for r := range A {
// 			A[r][c] = A[r][c] ^ 1
// 		}
// 	}
// 	for r := range A {
// 		if A[r][0] == 0 {
// 			toggleRow(r)
// 		}
// 	}
// 	for c := 1; c < col; c++ {
// 		cnt1 := 0
// 		for r := 0; r < row; r++ {
// 			if A[r][c] == 1 {
// 				cnt1++
// 			}
// 		}
// 		if cnt1 <= row/2 {
// 			toggleCol(c)
// 		}
// 	}
// 	var cal func(r int) int
// 	cal = func(r int) int {
// 		curRes := 0
// 		for c := col - 1; c >= 0; c-- {
// 			curRes += (1 << (col - c - 1)) * A[r][c]
// 		}
// 		return curRes
// 	}
// 	res := 0
// 	for r := range A {
// 		res += cal(r)
// 	}
// 	return res
// }
