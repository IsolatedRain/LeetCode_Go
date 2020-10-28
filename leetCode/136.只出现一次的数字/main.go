package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	res := nums[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		res ^= nums[i]
	}
	return res
}

func main() {
	p := []int{4, 1, 2, 1, 2, 4, 5}
	r := singleNumber(p)
	fmt.Printf("输入: %v\n输出: %v\n", p, r)
}
