package main

import "fmt"

func findTargetSumWays(nums []int, S int) int {
	n := len(nums)
	memo := map[[2]int]int{}
	var dfs func(curSum int, curID int) int
	dfs = func(curSum, curID int) int {
		if curID == n {
			if curSum == S {
				return 1
			}
			return 0
		}
		k := [2]int{curSum, curID}
		if v, ok := memo[k]; ok {
			return v
		}
		curRes := dfs(curSum+nums[curID], curID+1) + dfs(curSum-nums[curID], curID+1)
		memo[k] = curRes
		return curRes
	}
	return dfs(0, 0)
}

func main() {
	nums := []int{1, 1, 1, 1, 1}
	S := 3
	fmt.Println(findTargetSumWays(nums, S))
}
