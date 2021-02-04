package main

import "fmt"

func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	sum := 0
	for i := range nums {
		sum += nums[i]
		preSum[i+1] = sum
	}
	res := n + 1
	stack := []int{}
	for i := range preSum {
		// for len(stack) > 0 && preSum[stack[len(stack)-1]] >= preSum[i] {
		// 	stack = stack[:len(stack)-1]
		// }
		for len(stack) > 0 && preSum[i]-preSum[stack[0]] >= s {
			res = min(res, i-stack[0])
			stack = stack[1:]
		}
		stack = append(stack, i)
	}
	return res
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}
