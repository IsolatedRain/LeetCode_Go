// https://leetcode-cn.com/problems/find-the-most-competitive-subsequence/
package main

import "fmt"

func mostCompetitive(nums []int, k int) []int {
	stack := []int{}
	n := len(nums)
	for i := range nums {
		for len(stack) > 0 && stack[len(stack)-1] > nums[i] && len(stack)-1+n-i >= k {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return stack[:k]
}

func main() {
	nums := []int{84, 10, 71, 23, 66, 61, 62, 64, 34, 41, 80, 25, 91, 43, 4, 75, 65, 13, 37, 41, 46, 90, 55, 8, 85, 61, 95, 71}
	k := 24
	fmt.Println(mostCompetitive(nums, k))
}
