package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	curMin, curSecMax := nums[0], math.MaxInt64
	for _, num := range nums[1:] {
		if num > curSecMax {
			return true
		} else if num < curMin {
			curMin = num
		} else if curMin < num && num < curSecMax {
			curSecMax = num
		}
	}
	return false
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
}
