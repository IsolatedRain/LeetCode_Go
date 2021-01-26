package main

import "fmt"

func sortColors(nums []int) []int {
	L, R := 0, len(nums)-1
	sorted := true
	for L < R && sorted {
		sorted = false
		for i := L; i <= R; i++ {
			if nums[i] == 0 {
				sorted = true
				nums[i], nums[L] = nums[L], nums[i]
				L++
			} else if nums[i] == 2 {
				sorted = true
				nums[i], nums[R] = nums[R], nums[i]
				R--
			}
		}
	}
	return nums
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	fmt.Println(sortColors(nums))
}
