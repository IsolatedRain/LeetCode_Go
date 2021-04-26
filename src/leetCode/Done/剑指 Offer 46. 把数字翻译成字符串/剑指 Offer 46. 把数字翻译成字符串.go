// https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/
package main

import "fmt"

func translateNum(num int) int {
	if num < 10 {
		return 1
	}
	nums := []int{}
	p := 10
	for num > 0 {
		nums = append(nums, num%p)
		num /= p
	}
	L, R := 0, len(nums)-1
	for L < R {
		nums[L], nums[R] = nums[R], nums[L]
		L++
		R--
	}

	dp := make([]int, len(nums))
	dp[0], dp[1] = 1, 1
	if nums[1]+nums[0]*10 <= 25 {
		dp[1] = 2
	}

	for i := 2; i < len(nums); i++ {
		dp[i] = dp[i-1]
		if nums[i-1] != 0 && nums[i]+nums[i-1]*10 <= 25 {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(dp)-1]
}

func main() {
	fmt.Println(translateNum(506))
}
