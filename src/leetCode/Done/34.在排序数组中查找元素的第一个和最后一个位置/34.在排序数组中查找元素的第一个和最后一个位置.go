package main

import (
	"fmt"
	"sort"
)

func searchRange(nums []int, target int) []int {
	n := len(nums)
	L := sort.Search(n, func(i int) bool {
		return nums[i] >= target
	})
	if n == 0 || L == n || nums[L] != target {
		return []int{-1, -1}
	}
	R := sort.Search(n, func(i int) bool {
		return nums[i] > target
	})
	return []int{L, R - 1}
}

func main() {
	nums := []int{8, 8}
	target := 9
	r := searchRange(nums, target)
	fmt.Println(r)

}
