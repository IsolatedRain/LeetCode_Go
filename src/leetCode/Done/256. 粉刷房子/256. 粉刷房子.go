package main

import "fmt"

func minCost(costs [][]int) int {
	n := len(costs)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 3)
	}
	dp[0] = append([]int{}, costs[0]...)
	for i := 1; i < n; i++ {
		dp[i][0] = min(dp[i-1][1], dp[i-1][2]) + costs[i][0]
		dp[i][1] = min(dp[i-1][0], dp[i-1][2]) + costs[i][1]
		dp[i][2] = min(dp[i-1][0], dp[i-1][1]) + costs[i][2]
	}
	res := dp[n-1][0]
	for i := 1; i < 3; i++ {
		res = min(res, dp[n-1][i])
	}
	return res
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	costs := [][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}
	fmt.Println(minCost(costs))
}
