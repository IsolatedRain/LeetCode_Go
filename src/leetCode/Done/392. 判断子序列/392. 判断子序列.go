package main

import (
	"fmt"
)

func isSubsequence(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)
	if sLen == 0 {
		return true
	}
	if sLen > tLen {
		return false
	}
	if sLen == tLen {
		return s == t
	}
	dp := make([][26]int, tLen+1)
	for i := range dp {
		for j := 0; j < 26; j++ {
			dp[i][j] = -1
		}
	}
	dp[tLen-1][t[tLen-1]-'a'] = tLen - 1
	for r := tLen - 2; r > -1; r-- {
		for c := 0; c < 26; c++ {
			if dp[r+1][c] != -1 {
				dp[r][c] = dp[r+1][c]
			}
			dp[r][t[r]-'a'] = r
		}

	}
	r := 0
	for i := range s {
		if dp[r][s[i]-'a'] == -1 {
			return false
		}
		r = dp[r][s[i]-'a'] + 1

	}
	return true
}

func main() {
	s := "abcc"
	t := "ahbgadc"
	fmt.Println(isSubsequence(s, t))
}
