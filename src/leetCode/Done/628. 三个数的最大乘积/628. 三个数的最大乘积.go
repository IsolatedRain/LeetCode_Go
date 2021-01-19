package main

import "fmt"

func maximumProduct(nums []int) int {
	max0, max1, max2 := -1001, -1001, -1001
	min0, min1 := 1001, 1001
	for _, v := range nums {
		if v > max0 {
			max0, max1, max2 = v, max0, max1
		} else if v > max1 {
			max1, max2 = v, max1
		} else if v > max2 {
			max2 = v
		}
		if v < min0 {
			min0, min1 = v, min0
		} else if v < min1 {
			min1 = v
		}
	}
	a := min0 * min1 * max0
	b := max0 * max1 * max2
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(maximumProduct(nums))
}
