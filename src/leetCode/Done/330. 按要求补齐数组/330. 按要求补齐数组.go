package main

import "fmt"

// 1 -> k 连续， 则 1->k 的数，可以凑出 1 -> k+k-1 的连续数
func minPatches(nums []int, n int) int {
	count := 0
	numsLen := len(nums)
	i := 0
	missNum := 1
	for missNum <= n {
		if i < numsLen && nums[i] <= missNum {
			missNum += nums[i]
			i++
		} else {
			missNum = missNum + missNum - 1 + 1
			count++
		}
	}
	return count
}

func main() {
	nums := []int{1, 5, 10}
	n := 20
	r := minPatches(nums, n)
	fmt.Println(r)
}
