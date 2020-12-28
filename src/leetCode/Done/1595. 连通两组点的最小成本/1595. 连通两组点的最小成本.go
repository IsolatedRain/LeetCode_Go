package main

import (
	"fmt"
	"math"
)

func connectTwoGroups(cost [][]int) int {
	row := len(cost)
	col := len(cost[0])
	minCost := []int{}
	for c := 0; c < col; c++ {
		cur := cost[0][c]
		for r := 1; r < row; r++ {
			if cost[r][c] < cur {
				cur = min(cur, cost[r][c])
			}
		}
		minCost = append(minCost, cur)
	}
	memo := map[[2]int]int{}
	var f func(i, mask int) int
	f = func(i, mask int) int {
		if i == row {
			cur := 0
			for j := 0; j < col; j++ {
				if mask&(1<<j) == 0 {
					cur += minCost[j]
				}
			}
			return cur
		}
		k := [2]int{i, mask}
		if v, ok := memo[k]; ok {
			return v
		}
		res := math.MaxInt64
		for j := 0; j < col; j++ {
			curRes := cost[i][j] + f(i+1, mask|(1<<j))
			res = min(res, curRes)
		}
		memo[k] = res
		return res
	}
	return f(0, 0)
}

func min(x ...int) int {
	r := x[0]
	for _, v := range x[1:] {
		if v < r {
			r = v
		}
	}
	return r
}

func main() {
	cost := [][]int{{2, 5, 1}, {3, 4, 7}, {8, 1, 2}, {6, 2, 4}, {3, 8, 8}}
	r := connectTwoGroups(cost)
	fmt.Println(r)
}
