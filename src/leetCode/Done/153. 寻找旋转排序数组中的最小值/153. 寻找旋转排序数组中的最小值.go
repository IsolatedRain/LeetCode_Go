package main

import "fmt"

func findMin(nums []int) int {
	L, R := 0, len(nums)-1
	for L < R {
		mid := (L + R) >> 1
		if nums[mid] < nums[R] {
			R = mid
		} else {
			L = mid + 1
		}
	}
	return nums[L]
}

func main() {
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
}
