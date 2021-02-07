package main

import (
	"fmt"
)

func cherryPickup(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	var dp func(r, c1, c2 int) int
	memo := map[[3]int]int{}
	dp = func(r, c1, c2 int) int {
		if r == row-1 {
			if c1 != c2 {
				return grid[r][c1] + grid[r][c2]
			}
			return grid[r][c1]
		}
		k := [3]int{r, c1, c2}
		if v, ok := memo[k]; ok {
			return v
		}
		curRes := 0
		for nc1 := c1 - 1; nc1 < c1+2; nc1++ {
			for nc2 := c2 - 1; nc2 < c2+2; nc2++ {
				if nc1 >= 0 && nc2 >= 0 && nc1 < col && nc2 < col {
					curRes = max(curRes, dp(r+1, nc1, nc2))
				}
			}
		}
		if c1 == c2 {
			curRes += grid[r][c1]
		} else {
			curRes += grid[r][c1] + grid[r][c2]
		}
		memo[k] = curRes
		return curRes
	}
	return dp(0, 0, col-1)
}
func max(a ...int) int {
	r := a[0]
	for _, v := range a[1:] {
		if v > r {
			r = v
		}
	}
	return r
}

func main() {
	grid := [][]int{{3, 1, 1}, {2, 5, 1}, {1, 5, 5}, {2, 1, 1}}
	fmt.Println(cherryPickup(grid))
}
