package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	L, R := 0, len(s)-1
	for L < R {
		if !isAlpha(s[L]) && !isDigit(s[L]) {
			L++
			continue
		}
		if !isAlpha(s[R]) && !isDigit(s[R]) {
			R--
			continue
		}
		if s[L] != s[R] {
			return false
		}
		L++
		R--
	}
	return true
}

func isDigit(s byte) bool {
	if '0' <= s && s <= '9' {
		return true
	}
	return false
}
func isAlpha(s byte) bool {
	if 'a' <= s && s <= 'z' {
		return true
	}
	return false
}

func main() {
	fmt.Println(isPalindrome("0P"))
}
