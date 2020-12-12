package main

import "fmt"

func combine(n int, k int) (res [][]int) {
	arr := make([]int, k)
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if j == k {
			res = append(res, append([]int{}, arr...))
			return
		}
		for idx := i; idx < n+1; idx++ {
			arr[j] = idx
			dfs(idx+1, j+1)
		}
	}
	dfs(1, 0)
	return
}

func main() {
	n := 4
	k := 2
	r := combine(n, k)
	fmt.Println(r)
}
