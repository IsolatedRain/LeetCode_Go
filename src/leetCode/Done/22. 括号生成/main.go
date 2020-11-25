package main

import "fmt"

// 逐个添加合法的 ( 或 ), 当 ( == ) == n 的时候, 写入结果.
func generateParenthesis(n int) []string {
	var dfs func(s string, x, y int)
	res := []string{}
	dfs = func(curP string, leftP, rightP int) {
		if rightP == leftP && leftP == n {
			res = append(res, curP)
			return
		}
		if leftP < n {
			dfs(curP+"(", leftP+1, rightP)
		}
		if rightP < leftP && rightP <= n {
			dfs(curP+")", leftP, rightP+1)
		}
	}

	dfs("", 0, 0)
	return res
}

func main() {
	p := 5
	r := generateParenthesis(p)
	fmt.Printf("输入: %v\n输出:%v\n", p, r)
}
