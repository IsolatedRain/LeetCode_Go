package main

import "fmt"

func numDistinct(s string, t string) int {
	sLen, tLen := len(s), len(t)
	dp := make([][]int, tLen+1)
	for i := range dp {
		dp[i] = make([]int, sLen+1)
	}
	for i := range dp[0] {
		dp[0][i] = 1
	}
	for i := 1; i < tLen+1; i++ {
		for j := 1; j < sLen+1; j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[tLen][sLen]
}

// func numDistinct(s string, t string) int {
// 	var dfs func(a, b string) int
// 	count := func(s1 byte, s2 string) int {
// 		r := 0
// 		for i := range s2 {
// 			if s2[i] == s1 {
// 				r++
// 			}
// 		}
// 		return r
// 	}
// 	memo := map[[2]string]int{}
// 	dfs = func(a, b string) int {
// 		k := [2]string{a, b}
// 		if v, ok := memo[k]; ok {
// 			return v
// 		}
// 		if len(a) == 1 {
// 			return count(a[0], b)
// 		}
// 		if len(a) > len(b) || len(a) == 0 {
// 			return 0
// 		}
// 		r := 0
// 		for i := range b {
// 			if a[0] == b[i] {
// 				r += dfs(a[1:], b[i+1:])
// 			}
// 		}
// 		memo[k] = r
// 		return r
// 	}
// 	return dfs(t, s)
// }

func main() {
	s := "rabbbit"
	t := "rabbit"
	fmt.Println(numDistinct(s, t))
}
