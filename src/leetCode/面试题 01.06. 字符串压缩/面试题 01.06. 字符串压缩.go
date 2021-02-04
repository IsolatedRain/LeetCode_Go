package main

import (
	"fmt"
	"strconv"
)

func compressString(s string) string {
	s += " "
	res := ""
	curChar, count := s[0], 1
	for i := 1; i < len(s); i++ {
		if s[i] == curChar {
			count++
		} else {
			res += string(curChar) + strconv.Itoa(count)
			count = 1
			curChar = s[i]
		}
	}
	if len(res) >= len(s) {
		return s
	}
	return res
}

func main() {
	s := "aabcccccaaa"
	fmt.Println(compressString(s))
}
