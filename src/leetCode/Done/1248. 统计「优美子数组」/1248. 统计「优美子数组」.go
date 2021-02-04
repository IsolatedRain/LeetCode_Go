package main

import "fmt"

func numberOfSubarrays(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	sum, res := 0, 0
	for i := range nums {
		if nums[i]&1 == 1 {
			sum++
		}
		dp[sum]++
		if sum >= k {
			res += dp[sum-k]
		}
	}
	return res
}

func main() {
	nums := []int{1, 1, 2, 1, 1}
	k := 3
	fmt.Println(numberOfSubarrays(nums, k))
}
