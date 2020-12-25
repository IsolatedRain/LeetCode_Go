package main

import "fmt"

func integerReplacement(n int) int {
	memo := map[int]int{}
	var dfs func(n int) int
	dfs = func(n int) int {
		if n == 1 {
			return 0
		}
		if k, ok := memo[n]; ok {
			return k
		}
		curRes := 1
		if n&1 == 0 {
			curRes += dfs(n >> 1)
		} else {
			curRes += min(dfs(n+1), dfs(n-1))
		}
		memo[n] = curRes
		return memo[n]
	}
	return dfs(n)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	n := 1
	fmt.Println(integerReplacement(n))

}
