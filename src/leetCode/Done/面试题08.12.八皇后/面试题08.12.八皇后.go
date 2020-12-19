package main

import "fmt"

func solveNQueens(n int) [][]string {
	res := [][]string{}
	solves := make([][]byte, n)
	row := []byte{}
	for i := 0; i < n; i++ {
		row = append(row, '.')
	}
	for i := range solves {
		solves[i] = append([]byte{}, row...)
	}
	var dfs func(r, c, d, s int)
	dfs = func(r, c, d, s int) {
		if r == n {
			cur := []string{}
			for i := 0; i < n; i++ {
				cur = append(cur, string(solves[i]))
			}
			res = append(res, cur)
			return
		}
		uMask := c | d | s
		for i := 0; i < n; i++ {
			if 1<<i&uMask == 0 {
				curC := 1 << i
				solves[r][i] = 'Q'
				dfs(r+1, c|curC, (d|curC)<<1, (s|curC)>>1)
				solves[r][i] = '.'
			}
		}
	}
	dfs(0, 0, 0, 0)
	return res
}

func main() {
	n := 4
	r := solveNQueens(n)
	fmt.Println(r)
}
