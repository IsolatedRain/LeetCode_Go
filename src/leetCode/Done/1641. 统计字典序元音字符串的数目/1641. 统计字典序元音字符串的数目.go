package main

import (
	"fmt"
)

func countVowelStrings(n int) int {
	pre := make([]int, 6)
	for i := range pre {
		pre[i] = i
	}
	for i := 1; i < n; i++ {
		cur := make([]int, 6)
		for j := 1; j < 6; j++ {
			cur[j] = pre[j] + cur[j-1]
		}
		pre = cur
	}
	return pre[5]
}

func main() {
	fmt.Println(countVowelStrings(33))
}
