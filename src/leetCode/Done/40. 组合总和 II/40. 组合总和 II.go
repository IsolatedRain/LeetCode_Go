package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	n := len(candidates)
	res := [][]int{}
	path := []int{}
	var f func(curID int, curSum int)
	f = func(curID int, curSum int) {
		if curSum == target {
			res = append(res, append([]int{}, path...))
		}
		for i := curID; i < n; i++ {
			if candidates[i]+curSum > target {
				return
			}
			if i > curID && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			f(i+1, curSum+candidates[i])
			path = path[:len(path)-1]
		}
	}
	f(0, 0)
	return res
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	r := combinationSum2(candidates, target)
	fmt.Println(r)
}
