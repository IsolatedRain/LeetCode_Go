package main

import (
	"fmt"
	"math"
)

func find132pattern(nums []int) bool {
	stack := []int{}
	secMax := math.MinInt64
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		if nums[i] < secMax {
			return true
		}
		for len(stack) > 0 && stack[len(stack)-1] < nums[i] {
			secMax = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false
}

func main() {
	fmt.Println(find132pattern([]int{3, 1, 4, 2}))
}
