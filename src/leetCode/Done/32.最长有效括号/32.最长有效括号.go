package main

import "fmt"

func longestValidParentheses(s string) int {
	n := len(s)
	res := 0
	dp := make([]int, n)
	stack := []int{}
	for i := range s {
		if s[i] == '(' {
			dp[i] = 0
			stack = append(stack, i)
		} else if s[i] == ')' {
			if len(stack) == 0 {
				dp[i] = 0
			} else {
				left := stack[len(stack)-1]
				curSize := i - left + 1
				stack = stack[:len(stack)-1]
				if left > 0 {
					curSize += dp[left-1]
				}
				dp[i] = curSize
				if curSize > res {
					res = curSize
				}
			}
		}
	}
	// fmt.Println(dp)
	return res
}
func main() {
	s := ")()())"
	r := longestValidParentheses(s)
	fmt.Println(r)
}
