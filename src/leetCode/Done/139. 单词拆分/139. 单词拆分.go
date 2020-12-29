package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	wordMap := map[string]bool{}
	for _, w := range wordDict {
		wordMap[w] = true
	}
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i < n+1; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}

func main() {
	s := "applepenapple"
	wordDict := []string{"apple", "pen"}
	r := wordBreak(s, wordDict)
	fmt.Println(r)
}
