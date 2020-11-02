package main

import "fmt"

func runningSum(nums []int) []int {
	n := len(nums)
	res := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		res[i] = nums[i-1] + res[i-1]
	}
	return res[1:]

}

func main() {
	p := []int{1, 1, 1, 1, 1}
	fmt.Printf("输入: %v\n", p)
	r := runningSum(p)
	fmt.Printf("输出: %v\n", r)
}
