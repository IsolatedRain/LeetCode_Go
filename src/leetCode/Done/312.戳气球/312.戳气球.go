package main

import "fmt"

func maxCoins(nums []int) int {
	nums = append(nums, 1)
	nums = append([]int{1}, nums...)
	n := len(nums)
	dp := [][]int{}
	for range nums {
		dp = append(dp, make([]int, n))
	}
	for R := 2; R < n; R++ {
		for L := R - 1; L > -1; L-- {
			for k := L + 1; k < R; k++ {
				fmt.Println(L, R, k, nums[L]*nums[k]*nums[R])
				dp[L][R] = max(dp[L][R], dp[L][k]+nums[L]*nums[k]*nums[R]+dp[k][R])
			}
		}
	}
	return dp[0][n-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{3, 1, 5, 8}
	r := maxCoins(nums)
	fmt.Println(r)
}
