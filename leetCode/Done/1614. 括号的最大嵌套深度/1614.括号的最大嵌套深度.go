package main

import "fmt"

func maxDepth(s string) int {
	curDepth := 0
	res := 0
	for _, p := range s {
		// if p == '(' {
		// 	curDepth++
		// } else if p == ')' {
		// 	curDepth--
		// }
		switch p {
		case '(':
			curDepth++
		case ')':
			curDepth--
		}
		res = max(res, curDepth)
	}
	return res
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	// p := "(1+(2*3)+((8)/4))+1"
	p := "(1)+((2))+(((3)))"
	fmt.Printf("输入: %v\n", p)
	r := maxDepth(p)
	fmt.Printf("输出: %v\n", r)
}
