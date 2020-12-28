package main

import "fmt"

func waysToDistribute(n int, k int) int {
	dp := make([][]int, k+1)
	mod := 1000000007
	for i := 0; i < k+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i < n+1; i++ {
		dp[1][i] = 1
	}
	for i := 2; i < k+1; i++ {
		for j := i; j < n+1; j++ {
			dp[i][j] = (dp[i-1][j-1] + dp[i][j-1]*i) % mod
		}
	}
	return dp[k][n]
}

func main() {
	n := 20
	k := 5
	r := waysToDistribute(n, k)
	fmt.Println(r)
}
