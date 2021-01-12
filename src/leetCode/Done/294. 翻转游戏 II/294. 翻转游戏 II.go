package main

import "fmt"

func canWin(s string) bool {
	memo := map[string]bool{}
	var dfs func(curS string) bool
	dfs = func(curS string) bool {
		if v, ok := memo[curS]; ok {
			return v
		}
		if len(curS) < 2 {
			return false
		}
		n := len(curS)
		for i := 0; i < n-1; i++ {
			if s[i] == '+' && s[i+1] == '+' {
				if dfs(curS[:i] + curS[i+2:]) {
					memo[curS] = true
					return true
				}
			}
		}
		memo[curS] = false
		return false
	}
	return dfs(s)
}
func main() {
	s := "++-+-++++"
	fmt.Println(canWin(s))
}
