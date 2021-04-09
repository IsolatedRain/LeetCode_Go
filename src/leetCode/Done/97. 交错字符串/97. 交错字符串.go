package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)
	if n3 != n1+n2 {
		return false
	}
	dp := make([][]bool, n1+1)
	for i := range dp {
		dp[i] = make([]bool, n2+1)
	}
	dp[0][0] = true
	for i := 0; i < n1+1; i++ {
		for j := 0; j < n2+1; j++ {
			k := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || (dp[i-1][j] && s1[i-1] == s3[k])
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || (dp[i][j-1] && s2[j-1] == s3[k])
			}
		}
	}
	return dp[n1][n2]
}

func main() {
	fmt.Println(isInterleave("aabcc",
		"dbbca",
		"aadbbcbcac"))
}

// func isInterleave(s1 string, s2 string, s3 string) bool {
// 	n1, n2, n3 := len(s1), len(s2), len(s3)
// 	if n3 != n1+n2 {
// 		return false
// 	}
// 	var f func(int, int, int) bool
// 	f = func(i, j, k int) bool {
// 		if k == n3 {
// 			return true
// 		}
// 		if i < n1 && s1[i] == s3[k] {
// 			if f(i+1, j, k+1) {
// 				return true
// 			}
// 		}
// 		if j < n2 && s2[j] == s3[k] {
// 			if f(i, j+1, k+1) {
// 				return true
// 			}
// 		}
// 		return false
// 	}
// 	return f(0, 0, 0)
// }
