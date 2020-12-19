package main

import "fmt"

func totalNQueens(n int) int {
	if n <= 1 {
		return 1
	}
	if n <= 3 {
		return 0
	}
	res := 0
	var dfs func(r, c, d, s int)
	dfs = func(r, c, d, s int) {
		if r == n {
			res++
			return
		}
		uMask := c | d | s
		for i := 0; i < n; i++ {
			if 1<<i&uMask == 0 {
				curC := 1 << i
				dfs(r+1, c|curC, (curC|d)<<1, (curC|s)>>1)
			}
		}
	}
	dfs(0, 0, 0, 0)
	return res
}

func main() {
	n := 8
	r := totalNQueens(n)
	fmt.Println(r)
}
