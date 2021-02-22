package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	nums = append(nums, 0)
	res, count := 0, 0

	for i := range nums {
		if nums[i] == 1 {
			count++
		} else {
			res = max(res, count)
			count = 0
		}
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
}
