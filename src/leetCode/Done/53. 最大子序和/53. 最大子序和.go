package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	curSum := nums[0]
	res := nums[0]
	for i := 1; i < n; i++ {
		if curSum < 0 {
			curSum = nums[i]
		} else {
			curSum += nums[i]
		}
		res = max(curSum, res)
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
	// nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	nums := []int{0, -1, -1, -1, -5, 3, -2, 5}
	r := maxSubArray(nums)
	fmt.Println(r)
}
