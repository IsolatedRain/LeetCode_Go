package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	L, R := 0, 0
	n := len(nums)
	if n == 0 {
		return 0
	}
	for R < n {
		if nums[R] != nums[L] {
			L++
			nums[L] = nums[R]
		}
		R++
	}
	return L + 1
}

func main() {
	para := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// para := []int{1, 1, 2}
	// para := []int{}
	// para := []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	r := removeDuplicates(para)
	fmt.Printf("%v\n", r)

}
