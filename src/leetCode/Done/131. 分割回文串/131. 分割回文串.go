package main

import "fmt"

func partition(s string) [][]string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		for j := i; j >= 0; j-- {
			if s[i] == s[j] {
				if j-i < 2 {
					dp[i][j] = true
				} else {
					if dp[i-1][j+1] {
						dp[i][j] = true
					}
				}
			}
		}
	}
	res := [][]string{}
	var dfs func(L int, path []string)
	dfs = func(L int, path []string) {
		if L == n {
			res = append(res, append([]string{}, path...))
			return
		}
		for i := L + 1; i < n+1; i++ {
			if dp[i-1][L] {
				dfs(i, append(path, s[L:i]))
			}
		}
	}
	dfs(0, []string{})
	return res
}

func main() {
	fmt.Println(partition("aab"))
}
