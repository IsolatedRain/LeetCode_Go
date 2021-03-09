// https://www.codewars.com/kata/5277c8a221e209d3f6000b56/train/go
package main

import "fmt"

// ValidBraces 有效的括号
func ValidBraces(p string) bool {
	stack := []byte{}
	pairs := map[byte]byte{'}': '{', ')': '(', ']': '['}
	for i := range p {
		if p[i] == '[' || p[i] == '{' || p[i] == '(' {
			stack = append(stack, p[i])
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[p[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(ValidBraces("([{}])"), true)
}
