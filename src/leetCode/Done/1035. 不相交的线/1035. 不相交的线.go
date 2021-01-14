package main

import "fmt"

func maxUncrossedLines(A []int, B []int) int {
	aLen, bLen := len(A), len(B)
	dp := make([][]int, aLen+1)
	for i := range dp {
		dp[i] = make([]int, bLen+1)
	}
	for i := 0; i < aLen; i++ {
		for j := 0; j < bLen; j++ {
			if A[i] == B[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}

		}
	}
	return dp[aLen][bLen]
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	A := []int{2, 5, 1, 2, 5}
	B := []int{10, 5, 2, 1, 5, 2}
	fmt.Println(maxUncrossedLines(A, B))
}
