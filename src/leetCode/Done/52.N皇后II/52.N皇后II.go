package main

import "fmt"

func totalNQueens(n int) int {
	if n <= 1 {
		return 1
	}
	if n <= 3 {
		return 0
	}
	maxMask := (1 << n) - 1
	res := 0
	var dfs func(r, c, d, s int)
	dfs = func(r, c, d, s int) {
		if r == n {
			res++
			return
		}
		curMask := maxMask & ^(c | d | s)
		for curMask > 0 {
			nxtC := curMask & (-curMask)
			dfs(r+1, c|nxtC, (d|nxtC)>>1, (s|nxtC)<<1)
			curMask &= ^nxtC
		}
	}
	dfs(0, 0, 0, 0)
	return res
}

func main() {
	n := 12
	r := totalNQueens(n)
	fmt.Println(r)
}
