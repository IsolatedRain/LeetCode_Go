package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	// 去除两边的 空格
	s = strings.TrimSpace(s)
	n := len(s)
	j := -1
	// fmt.Printf("%v\n", n)
	for i := len(s) - 1; i > -1; i-- {
		if s[i] == ' ' {
			j = i
			break
		}
	}
	if j == -1 {
		return n
	}
	return n - j - 1
}

func main() {
	// para := " Hello World"
	para := "abc      "
	r := lengthOfLastWord(para)
	fmt.Printf("%v\n", r)

}
