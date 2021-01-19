package main

import "fmt"

func findDuplicate(nums []int) int {
	for i := range nums {
		if nums[i]-1 == i {
			continue
		}
		for nums[i]-1 != i {
			if nums[i] == nums[nums[i]-1] {
				return nums[i]
			}
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	return -1
}

func main() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(nums))
}
