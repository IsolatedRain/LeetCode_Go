package main

import (
	"fmt"
	"math"
)

func cherryPickup(grid [][]int) int {
	n := len(grid)
	var dp func(r1, c1, r2 int) int
	memo := map[[3]int]int{}
	dp = func(r1, c1, r2 int) int {
		c2 := r1 + c1 - r2
		if r1 == n || c1 == n || r2 == n || c2 == n || grid[r1][c1] == -1 || grid[r2][c2] == -1 {
			return -math.MaxInt16
		}
		if r1 == c1 && r1 == n-1 {
			return grid[r1][c1]
		}
		k := [3]int{r1, c1, r2}
		if v, ok := memo[k]; ok {
			return v
		}
		res := grid[r1][c1] + grid[r2][c2]
		if r1 == r2 && c1 == c2 {
			res /= 2
		}
		res += max(dp(r1, c1+1, r2+1), dp(r1+1, c1, r2+1), dp(r1+1, c1, r2), dp(r1, c1+1, r2))
		memo[k] = res
		return res
	}
	return max(0, dp(0, 0, 0))
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
	fmt.Println(cherryPickup([][]int{{0, 1, -1}, {1, 0, -1}, {1, 1, 1}}))
}
