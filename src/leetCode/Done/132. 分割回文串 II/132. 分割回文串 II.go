package main

import (
	"fmt"
	"math"
)

func minCut(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		for j := range dp[i] {
			dp[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
		}
	}

	cut := make([]int, n)
	for i := range cut {
		if !dp[0][i] {
			cut[i] = math.MaxInt64
			for j := 0; j < i; j++ {
				if dp[j+1][i] && cut[j]+1 < cut[i] {
					cut[i] = cut[j] + 1
				}
			}
		}
	}
	return cut[n-1]
}

func main() {
	s := "aab"
	fmt.Println(minCut(s))
}
