package main

import (
	"fmt"
)

// 单调递增的栈， 栈长 刚好是数组个数。
func validSubarrays(nums []int) int {
	res := 0
	stack := []int{}
	for _, num := range nums {
		for len(stack) > 0 && stack[len(stack)-1] > num {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
		// fmt.Println(stack, num)
		res += len(stack)
	}
	return res
}

func main() {
	nums := []int{3, 2, 1}
	// nums := []int{2, 2, 2}
	// nums := []int{1, 4, 2, 5, 3}
	r := validSubarrays(nums)
	fmt.Println(r)
}
