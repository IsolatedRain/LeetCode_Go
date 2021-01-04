package main

import (
	"fmt"
)

func largestSubarray(nums []int, k int) []int {
	n := len(nums)
	curNum := nums[0]
	maxNumID := 0
	for i := 1; i < n-k+1; i++ {
		if curNum < nums[i] {
			curNum = nums[i]
			maxNumID = i
		}
	}
	return nums[maxNumID : maxNumID+k]
}

func main() {
	nums := []int{1, 4, 5, 2, 3}
	k := 3
	r := largestSubarray(nums, k)
	fmt.Println(r)
}
