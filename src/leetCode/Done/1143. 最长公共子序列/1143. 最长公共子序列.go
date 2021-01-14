package main

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	t1Len, t2Len := len(text1), len(text2)
	if t1Len == 0 || t2Len == 0 {
		return 0
	}

	dp := make([][]int, t1Len)
	for i := range dp {
		dp[i] = make([]int, t2Len)
	}
	for i := range text2 {
		if text1[0] == text2[i] {
			dp[0][i] = 1
		} else {
			if i > 0 {
				dp[0][i] = dp[0][i-1]
			}
		}
	}
	for i := range text1 {
		if text2[0] == text1[i] {
			dp[i][0] = 1
		} else {
			if i > 0 {
				dp[i][0] = dp[i-1][0]
			}
		}
	}

	for i := 1; i < t1Len; i++ {
		for j := 1; j < t2Len; j++ {
			if text1[i] == text2[j] {
				if i > 0 && j > 0 {
					dp[i][j] = dp[i-1][j-1] + 1
				} else {
					dp[i][j] = 1
				}
			} else {
				if i == 0 || j == 0 {
					dp[i][j] = 0
				}
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[t1Len-1][t2Len-1]
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	text1 := "abcde"
	text2 := "ace"
	fmt.Println(longestCommonSubsequence(text1, text2))
}
