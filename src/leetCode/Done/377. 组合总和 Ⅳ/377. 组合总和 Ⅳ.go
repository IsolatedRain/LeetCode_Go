// https://leetcode-cn.com/problems/combination-sum-iv/
package main

import (
	"fmt"
	"sort"
)

func combinationSum4(nums []int, target int) int {
	sort.Ints(nums)
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i < target+1; i++ {
		for _, num := range nums {
			if num > i {
				break
			}
			dp[i] += dp[i-num]
		}
	}
	return dp[target]
}

func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
}
