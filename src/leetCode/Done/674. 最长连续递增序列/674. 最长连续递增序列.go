package main

import (
	"fmt"
)

func findLengthOfLCIS(nums []int) int {
	res := 1
	curNum := nums[0]
	curLen := 1
	for _, num := range nums[1:] {
		if num > curNum {
			curLen++
		} else {
			if curLen > res {
				res = curLen
			}
			curLen = 1
		}
		curNum = num
	}
	if curLen > res {
		res = curLen
	}
	return res
}

func main() {
	nums := []int{1, 3, 5, 4, 7}
	fmt.Println(findLengthOfLCIS(nums))
}
