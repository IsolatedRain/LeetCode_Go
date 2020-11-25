package main

import "fmt"

func divisorGame(N int) bool {
	dp := make([]bool, N+2, N+2) // N+2, 覆盖 N = 1的情况,以免越界.
	dp[1] = false
	dp[2] = true
	for i := 3; i <= N; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 && dp[j] == true {
				dp[i] = true
				break
			}
		}
	}
	return dp[N]
}

func main() {
	n := 3
	fmt.Printf("输出: %v\n", n)
	r := divisorGame(n)
	fmt.Printf("输出: %v\n", r)

}
