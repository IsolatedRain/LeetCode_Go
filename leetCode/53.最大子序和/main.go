package main

import (
	"fmt"
)

// 前缀和 如果小于0, 就舍弃前面的, 从当前开始看.
func maxSubArray(nums []int) int {
	preSum := nums[0]
	if nums[0] < 0 {
		preSum = 0
	}
	curMax := nums[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		preSum += nums[i]
		curMax = max(curMax, preSum)
		if preSum <= 0 {
			preSum = 0
		}
	}
	return curMax
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	p := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	r := maxSubArray(p)
	fmt.Printf("%v\n", r)
}
