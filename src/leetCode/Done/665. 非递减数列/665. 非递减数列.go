package main

import "fmt"

func checkPossibility(nums []int) bool {
	n := len(nums)
	var isInc func(i int) bool
	isInc = func(i int) bool {
		pre, cur := nums[i-1], nums[i]
		nums[i-1] = cur
		isValid := true
		for i := 1; i < n; i++ {
			if nums[i] < nums[i-1] {
				isValid = false
				break
			}
		}
		if isValid {
			return true
		}
		isValid = true
		nums[i-1], nums[i] = pre, pre
		for i := 1; i < n; i++ {
			if nums[i] < nums[i-1] {
				isValid = false
				break
			}
		}
		return isValid
	}

	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			if isInc(i) {
				return true
			}
			return false
		}
	}
	return false
}

func main() {
	fmt.Println(checkPossibility([]int{-1, 4, 2, 3}))
}
