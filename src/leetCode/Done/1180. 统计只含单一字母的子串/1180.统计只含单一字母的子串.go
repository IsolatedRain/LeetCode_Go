package main

import "fmt"

func countLetters(S string) int {
	S = S + " "
	res := 0
	count := 0
	var pre rune
	for _, v := range S {
		if v != pre {
			res += count * (count + 1) / 2
			count = 1
			pre = v
		} else {
			count++
		}
	}
	return res
}

func main() {
	s := "aaaaaaaaaa"
	r := countLetters(s)
	fmt.Printf("输入: %v\n输出: %v\n", s, r)
}
