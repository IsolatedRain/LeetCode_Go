package main

import "fmt"

func rob(nums []int) int {
	var f func(arr []int) int
	f = func(arr []int) int {
		fir, sec := 0, arr[0]
		for _, v := range arr[1:] {
			fir, sec = sec, max(sec, fir+v)
		}
		return sec
	}
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	return max(f(nums[1:]), f(nums[:len(nums)-1]))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	// fmt.Println(rob([]int{1, 2, 3, 1}))
	// fmt.Println(rob([]int{200, 3, 140, 20, 10}))
	fmt.Println(rob([]int{1, 3, 1, 3, 100}))
}
