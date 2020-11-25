package main

import (
	"fmt"
)

// DFS找出所有的可能组合
func findLexSmallestString(s string, a int, b int) string {
	n := len(s)
	nums := make(map[string]int, n)
	res := s
	var op1Add func(x string) string
	var op2Rotate func(x string) string
	var dfs func(x string)

	op1Add = func(x string) string {
		curRet := []byte(x)
		for i := 1; i < n; i += 2 {
			curRet[i] = byte('0' + (int(curRet[i]-'0')+a)%10)
		}
		return string(curRet)
	}
	op2Rotate = func(x string) string {
		return x[b:] + x[:b]
	}

	dfs = func(x string) {
		if _, ok := nums[x]; ok {
			return
		}
		nums[x] = 1
		if x < res {
			res = x
		}
		curX := op1Add(x)
		if _, ok := nums[curX]; !ok {
			dfs(curX)
		}
		curX = op2Rotate(x)
		if _, ok := nums[curX]; !ok {
			dfs(curX)
		}

	}
	dfs(s)
	return res
}

type para struct {
	s string
	a int
	b int
}

func main() {

	p := para{"5525", 9, 2}
	fmt.Printf("输入: %v\n", p)
	r := findLexSmallestString(p.s, p.a, p.b)
	fmt.Printf("输出: %v\n", r)
}
