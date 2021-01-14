package main

import "fmt"

func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	// dp[i][j] 表示 数字范围  i->j 时的最小花费
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	maxMoney := (1 + n) * n / 2
	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j < n+1; j++ {
			if j-i == 1 {
				dp[i][j] = i
			} else if j-i == 2 {
				dp[i][j] = i + 1
			} else {
				dp[i][j] = maxMoney
				for k := i; k < j; k++ {
					//  选择 k 位置： 大了： 花费 dp[i][k-1]+k
					//				 小了： 花费 dp[k+1][j]+k
					dp[i][j] = min(max(dp[i][k-1], dp[k+1][j])+k, dp[i][j])
				}
			}
		}
	}
	return dp[1][n]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	n := 6
	fmt.Println(getMoneyAmount(n))
}
