package main

import "fmt"

func numDistinct(s string, t string) int {
	var dfs func(a, b string) int
	count := func(s1 byte, s2 string) int {
		r := 0
		for i := range s2 {
			if s2[i] == s1 {
				r++
			}
		}
		return r
	}
	memo := map[[2]string]int{}
	dfs = func(a, b string) int {
		k := [2]string{a, b}
		if v, ok := memo[k]; ok {
			return v
		}
		if len(a) == 1 {
			return count(a[0], b)
		}
		if len(a) > len(b) || len(a) == 0 {
			return 0
		}
		r := 0
		for i := range b {
			if a[0] == b[i] {
				r += dfs(a[1:], b[i+1:])
			}
		}
		memo[k] = r
		return r
	}
	return dfs(t, s)
}

func main() {
	s := "rabbbit"
	t := "rabbit"
	fmt.Println(numDistinct(s, t))
}
