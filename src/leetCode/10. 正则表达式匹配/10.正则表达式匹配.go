package main

import "fmt"

func isMatch(s string, p string) bool {
	sLen := len(s)
	pLen := len(p)
	var match func(i, j int) bool
	match = func(i, j int) bool {
		if i < sLen && j < pLen {
			if p[j] == s[i] || p[j] == '.' {
				return true
			}
		}
		return false
	}
	var f func(i, j int) bool
	memo := make([][]bool, 0, sLen)
	for i := 0; i < sLen; i++ {
		memo = append(memo, make([]bool, pLen, pLen))
	}
	f = func(i, j int) bool {
		if i == sLen {
			return j == pLen
		}
		if i >= sLen || j >= pLen {
			return false
		}
		if memo[i][j] == true {
			return true
		}
		ast := false
		if j+1 < pLen && p[j+1] == '*' {
			ast = true
		}
		curMatch := match(i, j)
		var match bool

		if ast {
			if curMatch {
				match = f(i+1, j+2) || f(i+1, j) || f(i, j+2)
			} else {
				match = f(i, j+2)
			}
		} else {
			match = curMatch && f(i+1, j+1)
		}
		memo[i][j] = match
		return match
	}
	res := f(0, 0)

	return res
}

func main() {
	s := "mississippi"
	p := "mis*is*p*."
	fmt.Printf("输入: %v\n%v\n", s, p)
	r := isMatch(s, p)
	fmt.Printf("输出: %v\n", r)
}
