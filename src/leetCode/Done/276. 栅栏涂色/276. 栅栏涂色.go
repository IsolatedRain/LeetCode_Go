package main

import "fmt"

func numWays(n int, k int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return k
	}
	dp := make([]int, n)
	dp[0] = k
	same, diff := k, k
	for i := 1; i < n; i++ {
		same, diff = diff, dp[i-1]*(k-1)
		dp[i] = same + diff
	}
	return dp[n-1]
}

func main() {
	n := 5
	k := 10
	fmt.Println(numWays(n, k))
}
