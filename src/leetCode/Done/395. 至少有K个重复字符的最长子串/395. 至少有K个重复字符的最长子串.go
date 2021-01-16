package main

import (
	"fmt"
)

func longestSubstring(s string, k int) int {
	var count func(curS string) byte
	count = func(curS string) (b byte) {
		cnt := map[byte]int{}
		for i := range curS {
			cnt[curS[i]]++
		}
		for c, v := range cnt {
			if v < k {
				return c
			}
		}
		return b
	}
	var dfs func(curS string) int
	dfs = func(curS string) int {
		if len(curS) < k {
			return 0
		}
		c := count(curS)
		if c == byte(0) {
			return len(curS)
		}
		i := 0
		for i < len(curS) {
			if curS[i] == c {
				break
			}
			i++
		}
		left := dfs(curS[:i])
		right := dfs(curS[i+1:])
		return max(left, right)
	}
	return dfs(s)
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	s := "abaccbabc"
	k := 3
	fmt.Println(longestSubstring(s, k))
}
