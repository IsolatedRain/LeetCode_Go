package leetcode

import (
	"strings"
)

// 逐对替换, 直到为空
func isValid(s string) bool {
	for {
		old := s
		s = strings.Replace(s, "()", "", -1)
		s = strings.Replace(s, "{}", "", -1)
		s = strings.Replace(s, "[]", "", -1)
		if s == "" {
			return true
		}
		if s == old {
			return false
		}
		old = s
	}
}

// 栈, 逐对抵消 [] () {}
// func isValid(s string) bool {
// 	n := len(s)
// 	if n&1 == 1 {
// 		return false
// 	}
// 	if n == 0 {
// 		return true
// 	}
// 	pairs := map[byte]byte{
// 		')': '(',
// 		']': '[',
// 		'}': '{',
// 	}
// 	stack := []byte{}
// 	for i := 0; i < n; i++ {
// 		v, _ := pairs[s[i]]
// 		if v > 0 {
// 			if len(stack) == 0 || stack[len(stack)-1] != v {
// 				return false
// 			}
// 			stack = stack[:len(stack)-1]
// 		} else {
// 			stack = append(stack, s[i])
// 		}
// 	}
// 	return len(stack) == 0
// }
