package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	for _, v := range nums {
		nums[abs(v)-1] = -abs(nums[abs(v)-1])
	}
	res := []int{}
	for i, v := range nums {
		if v > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(nums))
}
