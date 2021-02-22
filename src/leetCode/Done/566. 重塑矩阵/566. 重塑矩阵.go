package main

import "fmt"

func matrixReshape(nums [][]int, r int, c int) [][]int {
	row, col := len(nums), len(nums[0])
	if r*c != row*col {
		return nums
	}
	res := make([][]int, r)
	for i := range res {
		res[i] = make([]int, c)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			v := i*c + j
			ni, nj := v/col, v%col
			res[i][j] = nums[ni][nj]
		}
	}
	return res
}

func main() {
	fmt.Println(matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4))
}
