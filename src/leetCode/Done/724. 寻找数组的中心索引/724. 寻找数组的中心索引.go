package main

import "fmt"

func pivotIndex(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	preSum := []int{0}
	preSum[0] = 0
	for i := range nums {
		preSum = append(preSum, preSum[len(preSum)-1]+nums[i])
	}
	preSum = append(preSum, 0)
	for i := 1; i < n+1; i++ {
		left := preSum[i-1]
		right := preSum[n] - preSum[i]
		if left == right {
			return i - 1
		}
	}

	return -1
}

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	// fmt.Println(pivotIndex([]int{-1, -1, -1, 0, 1, 1}))
}
