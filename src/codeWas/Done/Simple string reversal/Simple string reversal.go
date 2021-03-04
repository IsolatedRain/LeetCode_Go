//https://www.codewars.com/kata/5a71939d373c2e634200008e/train/go
package main

import (
	"fmt"
	"strings"
)

// Solve 翻转字符串， 保留原空格位置。
func Solve(s string) string {
	s1 := strings.Replace(s, " ", "", -1)
	j := len(s1) - 1
	r := ""
	for i := range s {
		if s[i] == ' ' {
			r += " "
		} else {
			r += string(s1[j])
			j--
		}
	}
	return r
}

func main() {
	fmt.Println(Solve("your code rocks") == "skco redo cruoy")
}
