package main

import "fmt"

func massage(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[n-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	// nums := []int{2, 1, 4, 5, 3, 1, 1, 3}
	nums := []int{2, 1, 1, 2}
	fmt.Println(massage(nums))
}
