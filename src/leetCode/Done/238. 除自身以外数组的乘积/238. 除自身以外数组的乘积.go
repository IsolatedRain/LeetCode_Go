package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	left := make([]int, n)
	cur := 1
	for i := 0; i < n; i++ {
		left[i] = cur * nums[i]
		cur = left[i]
	}
	res := make([]int, n)
	cur = 1
	for i := n - 1; i >= 1; i-- {
		res[i] = cur * left[i-1]
		cur *= nums[i]
	}
	res[0] = cur
	return res
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
