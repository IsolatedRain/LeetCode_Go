package main

import "fmt"

func numDecodings(s string) int {
	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	}
	prePre, pre, cur := 1, 1, 1
	for i := 2; i < n+1; i++ {
		cur = 0
		if s[i-1] == '0' {
			if s[i-2] != '1' && s[i-2] != '2' {
				return 0
			}
			cur = prePre
		} else if (s[i-2] == '0') || (s[i-1]-'0'+(s[i-2]-'0')*10) > 26 {
			cur = pre
		} else {
			cur = pre + prePre
		}
		prePre, pre = pre, cur
	}
	return cur
}

func main() {
	// s := "1"
	s := "2101"
	fmt.Println(numDecodings(s))
}
