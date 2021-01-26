package main

import "fmt"

func lengthOfLongestSubstringTwoDistinct(s string) int {
	stack := []int{}
	res := 0
	idx := map[byte]int{}
	for i := range s {
		idx[s[i]] = i
		stack = append(stack, i)
		for len(idx) > 2 {
			for j := range stack {
				t := s[stack[j]]
				if idx[t] == stack[0] {
					delete(idx, t)
				}
				stack = stack[j+1:]
				break
			}
		}
		res = max(res, i-stack[0]+1)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	s := "ccaabbb"
	fmt.Println(lengthOfLongestSubstringTwoDistinct(s))
}
