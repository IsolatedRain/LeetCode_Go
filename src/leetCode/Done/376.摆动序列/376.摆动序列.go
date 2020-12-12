package main

import "fmt"

func wiggleMaxLength(nums []int) int {
	up, down := 1, 1
	if len(nums) < 2 {
		return len(nums)
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}
	}
	if up > down {
		return up
	}
	return down
}

func main() {
	// nums := []int{1, 17, 5, 10, 13, 15, 15, 10, 5, 16, 8}
	nums := []int{1, 2, 3, 4, 4, 4, 5, 5, 5, 6, 5}
	r := wiggleMaxLength(nums)
	fmt.Println(r)

}
