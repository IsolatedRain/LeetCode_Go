// https://www.codewars.com/kata/546d5028ddbcbd4b8d001254/train/go
package main

import "fmt"

// Partitions 组成数字的方法数
func Partitions(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i < n+1; i++ {
		for j := i; j < n+1; j++ {
			dp[j] += dp[j-i]
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(Partitions(10) == 42)
}
