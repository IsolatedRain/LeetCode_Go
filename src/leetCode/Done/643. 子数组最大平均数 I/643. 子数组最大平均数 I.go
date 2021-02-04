package main

import (
	"fmt"
	"math"
)

func findMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	preSum := []int{0}
	for i := range nums {
		preSum = append(preSum, preSum[len(preSum)-1]+nums[i])
	}
	L, R := 0, k
	res := math.MaxFloat64 * -1
	for R < n+1 {
		res = max(res, float64(preSum[R]-preSum[L])/float64(k))
		R++
		L++
	}
	return res
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x

}
func max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	fmt.Println(findMaxAverage(nums, k))
}
