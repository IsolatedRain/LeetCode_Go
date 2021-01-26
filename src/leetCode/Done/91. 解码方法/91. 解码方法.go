package main

import "fmt"

func numDecodings(s string) int {
	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i < n+1; i++ {
		if s[i-1] == '0' {
			if s[i-2] != '1' && s[i-2] != '2' {
				return 0
			}
			dp[i] = dp[i-2]
		} else if s[i-2] != '1' && s[i-2] != '2' {
			dp[i] = dp[i-1]
		} else if (s[i-1] == '7' || s[i-1] == '8' || s[i-1] == '9') && s[i-2] == '2' {
			dp[i] = dp[i-1]
		} else {
			dp[i] = dp[i-1] + dp[i-2]
		}
	}
	return dp[n]
}

func main() {
	s := "226"
	fmt.Println(numDecodings(s))
}
