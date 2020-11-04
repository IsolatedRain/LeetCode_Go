package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var res [][]int
	n := len(nums)
	for i := 0; i < n-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		L := i + 1
		R := n - 1
		for L < R {
			curSum := nums[i] + nums[L] + nums[R]
			if curSum == 0 {
				res = append(res, []int{nums[i], nums[L], nums[R]})
				L++
				R--
				for L < R && nums[L] == nums[L-1] {
					L++
				}
				for L < R && nums[R] == nums[R+1] {
					R--
				}
			} else if curSum > 0 {
				R--
			} else if curSum < 0 {
				L++
			}
		}
	}
	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Printf("输入: %v\n", nums)
	r := threeSum(nums)
	fmt.Printf("输出: %v\n", r)
}
