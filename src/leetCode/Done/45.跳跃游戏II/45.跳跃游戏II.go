package main

import "fmt"

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	steps, lastMax, curMax, n := 1, nums[0], nums[0], len(nums)
	if nums[0] >= n-1 {
		return 1
	}
	for i := 1; i < n; i++ {
		curMax = max(curMax, i+nums[i])
		if curMax >= n-1 {
			return steps + 1
		}
		if i == lastMax {
			lastMax = curMax
			steps++
		}
	}
	return steps
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	// nums := []int{2, 3, 1, 1, 4}
	nums := []int{1, 2}
	r := jump(nums)
	fmt.Println(r)
}
