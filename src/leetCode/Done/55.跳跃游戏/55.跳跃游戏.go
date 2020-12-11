package main

import "fmt"

func canJump(nums []int) bool {
	curMax := nums[0]
	tar := len(nums) - 1
	for i, v := range nums {
		if curMax >= tar {
			return true
		}
		if i <= curMax {
			if i+v > curMax {
				curMax = i + v
			}
		}
	}
	return false
}

func main() {
	// nums := []int{2, 3, 1, 1, 4}
	nums := []int{3, 2, 1, 0, 4}
	r := canJump(nums)
	fmt.Println(r)
}
