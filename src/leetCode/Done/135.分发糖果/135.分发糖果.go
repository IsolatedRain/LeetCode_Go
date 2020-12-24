package main

import (
	"fmt"
)

func candy(ratings []int) int {
	n := len(ratings)
	left := make([]int, n)
	left[0] = 1
	right := make([]int, n)
	right[n-1] = 1
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	for i := n - 2; i > -1; i-- {
		if ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		} else {
			right[i] = 1
		}
	}
	res := 0
	for i := range left {
		res += max(left[i], right[i])
	}
	return res
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	ratings := []int{1, 2, 2}
	r := candy(ratings)
	fmt.Println(r)
}
