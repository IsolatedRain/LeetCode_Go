package main

import "fmt"

func findMin(nums []int) int {
	L, R := 0, len(nums)-1
	for L < R {
		mid := (L + R) >> 1
		if nums[mid] < nums[R] {
			R = mid
		} else if nums[mid] > nums[R] {
			L = mid + 1
		} else {
			R--
		}
	}
	return nums[L]
}

func main() {
	fmt.Println(findMin([]int{2, 2, 2, 0, 1}))
}
