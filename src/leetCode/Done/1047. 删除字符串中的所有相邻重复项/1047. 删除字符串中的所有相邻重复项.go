// https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string/
package main

import "fmt"

func removeDuplicates(S string) string {
	stack := []byte{}
	for i := range S {
		if len(stack) > 0 && stack[len(stack)-1] == S[i] {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, S[i])
	}
	return string(stack)
}

func main() {
	fmt.Println(removeDuplicates("abbaca"))
}
