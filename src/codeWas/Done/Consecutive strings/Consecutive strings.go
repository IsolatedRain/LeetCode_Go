// https://www.codewars.com/kata/56a5d994ac971f1ac500003e/train/go
package main

import (
	"fmt"
	"strings"
)

// LongestConsec k个连续字符串组成的最长字符
func LongestConsec(strarr []string, k int) string {
	fmt.Println(strarr, k)
	if len(strarr) == 0 || k <= 0 || len(strarr) < k {
		return ""
	}
	if k >= len(strarr) {
		return strings.Join(strarr, "")
	}
	preSum := make([]int, len(strarr)+1)
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + len(strarr[i-1])
	}
	curMax := 0
	L, R := -1, -1
	for i := k; i < len(preSum); i++ {
		if preSum[i]-preSum[i-k] > curMax {
			curMax = preSum[i] - preSum[i-k]
			R = i
			L = i - k
		}
	}
	return strings.Join(strarr[L:R], "")
}

func main() {
	s := []string{"it", "wkppv", "ixoyx", "3452", "zzzzzzzzzzzz"}
	k := 15
	// fmt.Println(LongestConsec([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2) == "abigailtheta")
	fmt.Println(LongestConsec(s, k))
}
