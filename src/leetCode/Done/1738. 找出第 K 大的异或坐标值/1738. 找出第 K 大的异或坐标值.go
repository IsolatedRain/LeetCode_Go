package main

import (
	"fmt"
	"sort"
)

func kthLargestValue(matrix [][]int, k int) int {
	row, col := len(matrix), len(matrix[0])
	dp := make([][]int, row)
	for i := range dp {
		dp[i] = make([]int, col)
	}
	values := []int{matrix[0][0]}
	dp[0][0] = matrix[0][0]
	for r := 1; r < row; r++ {
		dp[r][0] = matrix[r][0] ^ dp[r-1][0]
		values = append(values, dp[r][0])
	}
	for c := 1; c < col; c++ {
		dp[0][c] = matrix[0][c] ^ dp[0][c-1]
		values = append(values, dp[0][c])
	}
	for r := 1; r < row; r++ {
		for c := 1; c < col; c++ {
			dp[r][c] = dp[r-1][c] ^ dp[r][c-1] ^ dp[r-1][c-1] ^ matrix[r][c]
			values = append(values, dp[r][c])
		}
	}
	sort.Slice(values, func(i, j int) bool { return values[i] > values[j] })
	return values[k-1]
}

func main() {
	matrix := [][]int{{5, 2}, {1, 6}}
	k := 1
	fmt.Println(kthLargestValue(matrix, k))
}
