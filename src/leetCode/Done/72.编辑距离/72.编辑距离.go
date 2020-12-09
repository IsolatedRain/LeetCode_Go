package main

import "fmt"

// 数组重复利用
func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	last := make([]int, n2+1)
	cur := make([]int, n2+1)
	for i := 0; i < n2+1; i++ {
		last[i] = i
	}
	for i := 1; i < n1+1; i++ {
		cur[0] = i
		for j := 1; j < n2+1; j++ {
			if word1[i-1] == word2[j-1] {
				cur[j] = last[j-1]
			} else {
				cur[j] = min(cur[j-1], last[j], last[j-1]) + 1
			}
		}
		last, cur = cur, last
	}
	return last[n2]
}

func min(x ...int) int {
	r := x[0]
	for _, v := range x {
		if r > v {
			r = v
		}
	}
	return r
}

func main() {
	word1 := "intention"
	word2 := "execution"
	// word1 := "horse"
	// word2 := "ros"
	r := minDistance(word1, word2)
	fmt.Println(r)
}

// DP数组
// func minDistance(word1 string, word2 string) int {
// 	n1 := len(word1)
// 	n2 := len(word2)
// 	dp := make([][]int, n1+1)
// 	for i := range dp {
// 		dp[i] = make([]int, n2+1)
// 	}
// 	for i := 0; i < n2+1; i++ {
// 		dp[0][i] = i
// 	}
// 	for i := 0; i < n1+1; i++ {
// 		dp[i][0] = i
// 	}
// 	for i := 1; i < n1+1; i++ {
// 		for j := 1; j < n2+1; j++ {
// 			if word1[i-1] == word2[j-1] {
// 				dp[i][j] = dp[i-1][j-1]
// 			} else {
// 				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
// 			}
// 		}
// 	}
// 	return dp[n1][n2]
// }
