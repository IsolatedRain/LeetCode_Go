package main

import "fmt"

func trap(height []int) int {
	n := len(height)
	left := make([]int, n, n)
	right := make([]int, n, n)
	left[0] = height[0]
	right[n-1] = height[n-1]
	for i := 1; i < n; i++ {
		left[i] = max(height[i], left[i-1])
	}
	for i := n - 2; i >= 0; i-- {
		right[i] = max(height[i], right[i+1])
	}
	res := 0
	for i := 1; i < n-1; i++ {
		minH := min(left[i], right[i])
		if minH > height[i] {
			res += minH - height[i]
		}
	}
	return res
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
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	r := trap(height)
	fmt.Println(r)
}
