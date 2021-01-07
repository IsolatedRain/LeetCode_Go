package main

import "fmt"

func findJudge(N int, trust [][]int) int {
	d := make(map[int][]int, N+1)
	for i := 1; i < N+1; i++ {
		d[i] = make([]int, 2)
	}
	for _, t := range trust {
		d[t[0]][1]++
		d[t[1]][0]++
	}
	res := -1
	count := 0
	for k, v := range d {
		if v[0] == N-1 && v[1] == 0 {
			count++
			res = k
		}
	}
	if count == 1 {
		return res
	}
	return -1
}

func main() {
	N := 4
	trust := [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}
	fmt.Println(findJudge(N, trust))
}
