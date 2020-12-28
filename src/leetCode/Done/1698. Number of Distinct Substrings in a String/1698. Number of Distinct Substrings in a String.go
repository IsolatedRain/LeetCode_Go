package main

import (
	"fmt"
	"strings"
)

func countDistinct(s string) int {
	n := len(s)
	dp := make([]int, n)
	mark := make([]int, 26)
	dp[0] = 1
	mark[s[0]-'a'] = 1

	for i := 1; i < n; i++ {
		if mark[s[i]-'a'] == 0 {
			dp[i] = dp[i-1] + i + 1
			mark[s[i]-'a'] = 1
			continue
		}
		diffID := -1
		for j := i; j > -1; j-- {
			if !strings.Contains(s[:i], s[j:i+1]) {
				diffID = j + 1
				break
			}
		}
		dp[i] = dp[i-1] + diffID
	}
	return dp[len(dp)-1]
}

func main() {
	// s := "aabbaba"
	s := "abcdefg"
	// s := "hlq"
	// s := "aaa"
	// s := "aaaa"
	r := countDistinct(s)
	fmt.Println(r)
}

// 暴力
// func countDistinct(s string) int {
// 	n := len(s)
// 	mapSet := map[string]struct{}{}
// 	for i := 0; i < n; i++ {
// 		for j := n; j > i; j-- {
// 			if _, ok := mapSet[s[i:j]]; ok {
// 				break
// 			}
// 			mapSet[s[i:j]] = struct{}{}
// 		}
// 	}
// 	return len(mapSet)
// }
