package main

import "fmt"

func moveZeroes(nums []int) []int {
	L, R, n := 0, 1, len(nums)
	for R < n {
		fmt.Printf("%v %v %v\n", L, R, nums)
		if nums[L] != 0 {
			L++
		} else if nums[R] != 0 {
			nums[L], nums[R] = nums[R], nums[L]
			L++
		}
		R++
	}
	return nums
}

func main() {
	p := []int{0, 1, 0, 3, 12}
	r := moveZeroes(p)
	fmt.Printf("输入%v\n 输出:%v\n", p, r)
}
