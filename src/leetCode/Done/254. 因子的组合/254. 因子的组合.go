package main

import (
	"fmt"
)

func getFactors(n int) [][]int {
	res := [][]int{}
	var dfs func(curFac, curN int, factors []int)
	dfs = func(curFac, curN int, factors []int) {
		fmt.Println(curFac, curN, factors)
		if curN == 1 {
			if len(factors) > 1 {
				res = append(res, append([]int{}, factors...))
			}
			return
		}
		for f := curFac; f < curN+1; f++ {
			if curN%f == 0 {
				dfs(f, curN/f, append(factors, f))
			}
		}
	}
	dfs(2, n, []int{})
	return res
}

func main() {
	fmt.Println(getFactors(1024))
}
