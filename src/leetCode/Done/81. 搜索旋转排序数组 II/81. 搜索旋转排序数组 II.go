package main

import "fmt"

func search(nums []int, target int) bool {
	L, R := 0, len(nums)-1
	for L <= R {
		mid := (L + R) >> 1
		if nums[mid] == target {
			return true
		} else {
			// 左边有序
			if nums[mid] > nums[L] {
				if nums[mid] > target && target >= nums[L] {
					R = mid
				} else {
					L = mid + 1
				}
				// 右边有序
			} else if nums[mid] < nums[L] {
				if nums[mid] <= target && target <= nums[R] {
					L = mid
				} else {
					R = mid - 1
				}
			} else {
				L++
			}
		}
	}
	return false
}

func main() {
	nums := []int{2, 5, 6, 0, 0, 1, 1, 1}
	target := 2
	fmt.Println(search(nums, target))
}
