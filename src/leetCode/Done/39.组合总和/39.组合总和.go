package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	n := len(candidates)
	res := [][]int{}
	path := []int{}
	var dfs func(curID, curSum int)
	dfs = func(curID, curSum int) {
		if curSum == target {
			// 浅拷贝，slice会被后续修改， 需要新复制一份。
			tmp := append([]int{}, path...)
			res = append(res, tmp)
			return
		}
		for i := curID; i < n; i++ {
			if curSum+candidates[i] > target {
				break
			}
			path = append(path, candidates[i])
			dfs(i, curSum+candidates[i])
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return res
}

func main() {
	candidates := []int{3, 2, 5}
	target := 8
	r := combinationSum(candidates, target)
	fmt.Println(r)
}
