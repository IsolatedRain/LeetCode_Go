package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) (res [][]int) {
	n := len(nums)
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	curSub := []int{}
	var dfs func(int)
	dfs = func(idx int) {
		res = append(res, append([]int{}, curSub...))
		for i := idx; i < n; i++ {
			if i > idx && nums[i] == nums[i-1] {
				continue
			}
			curSub = append(curSub, nums[i])
			dfs(i + 1)
			curSub = curSub[:len(curSub)-1]
		}
	}
	dfs(0)
	return res
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
