package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	stack := []int{}
	res := []int{}
	for i := 0; i < n; i++ {
		if len(stack) == 0 {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && nums[i] >= nums[stack[len(stack)-1]] {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, i)
			if i-stack[0] >= k {
				stack = stack[1:]
			}
		}
		if i >= k-1 {
			res = append(res, nums[stack[0]])
		}
	}
	return res
}

func main() {
	nums := []int{3, 1, -1, -3, 5, 3, 6, 7}
	k := 1
	r := maxSlidingWindow(nums, k)
	fmt.Println(r)
}
