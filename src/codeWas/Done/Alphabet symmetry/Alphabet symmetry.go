// https://www.codewars.com/kata/59d9ff9f7905dfeed50000b0/train/go
package main

import (
	"fmt"
	"strings"
)

func solve(slice []string) []int {
	cal := func(s string) int {
		s = strings.ToLower(s)
		count := 0
		for i := range s {
			if int(s[i]-'a') == i {
				count++
			}
		}
		return count
	}
	res := make([]int, len(slice))
	for i, s := range slice {
		res[i] = cal(s)
	}
	return res
}

func main() {
	fmt.Println(solve([]string{"encode", "abc", "xyzD", "ABmD"}), []int{1, 3, 1, 3})
}
