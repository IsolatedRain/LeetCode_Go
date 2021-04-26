package main

import (
	"fmt"
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	n := len(nums)
	dp := make([]int, n)
	maxSize, maxID := -1, -1
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > maxSize {
			maxSize = dp[i]
			maxID = i
		}
	}
	res := []int{}
	curNum := nums[maxID]
	curSize := maxSize
	for i := maxID; i >= 0; i-- {
		if dp[i] == curSize && curNum%nums[i] == 0 {
			curSize--
			curNum = nums[i]
			res = append(res, curNum)
		}
	}
	return res
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(largestDivisibleSubset([]int{1, 2, 4, 8}))
}
