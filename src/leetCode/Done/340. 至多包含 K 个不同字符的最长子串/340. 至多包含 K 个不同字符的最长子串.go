package main

import "fmt"

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	stack := []int{}
	curChars := map[byte]int{}
	res := 0
	for i := range s {
		curChars[s[i]] = i
		stack = append(stack, i)
		if len(curChars) <= k {
			res = max(res, i-stack[0]+1)
		} else {
			for j := range stack {
				if curChars[s[stack[j]]] == stack[j] {
					delete(curChars, s[stack[j]])
					stack = stack[j+1:]
					break
				}
			}
		}
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
	s := "eceba"
	k := 2
	fmt.Println(lengthOfLongestSubstringKDistinct(s, k))
}
