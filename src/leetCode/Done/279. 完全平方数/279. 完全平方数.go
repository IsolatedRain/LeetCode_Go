package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i < n+1; i++ {
		if isSquare(i) {
			dp[i] = 1
			continue
		}
		dp[i] = 4
		for j := 1; j*j < i; j++ {
			dp[i] = min(dp[i], dp[j*j]+dp[i-j*j])
			if dp[i] == 2 {
				break
			}
		}
	}
	return dp[n]
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func isSquare(x int) bool {
	sqrt := int(math.Sqrt(float64(x)))
	if sqrt*sqrt == x {
		return true
	}
	return false
}

func main() {
	fmt.Println(numSquares(99))
}
