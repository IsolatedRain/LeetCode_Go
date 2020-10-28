package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	L := 0
	R := len(nums) - 1
	for L <= R {
		mid := ((R - L) >> 1) + L
		v := nums[mid]
		if v < target {
			L = mid + 1
		} else {
			R = mid - 1
		}

	}
	return L
}

func main() {
	p := []int{1, 3, 5, 6}
	tar := 6
	r := searchInsert(p, tar)
	fmt.Printf("%v\n", r)

}
