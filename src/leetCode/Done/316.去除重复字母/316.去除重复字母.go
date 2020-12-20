package main

import (
	"fmt"
)

// 栈顶元素 大于当前， 且计数>0， 则弹出。
func removeDuplicateLetters(s string) string {
	count := map[byte]int{}
	add := map[byte]int{}
	for i := range s {
		count[s[i]]++
	}
	res := []byte{}
	for i := range s {
		if add[s[i]] > 0 {
			count[s[i]]--
			continue
		}
		for len(res) > 0 && res[len(res)-1] > s[i] && count[res[len(res)-1]] > 0 {
			add[res[len(res)-1]]--
			res = res[:len(res)-1]
		}
		res = append(res, s[i])
		count[s[i]]--
		add[s[i]]++
	}
	return string(res)
}

func main() {
	s := "bbcaac"
	// s := "adbdc"

	// s := "cbacdcbc"
	// s := "cdadabcc"
	r := removeDuplicateLetters(s)
	fmt.Println(r)
}
