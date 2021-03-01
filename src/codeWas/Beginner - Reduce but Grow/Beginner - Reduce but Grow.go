package main

import (
	"fmt"
	"strconv"
)

// Grow 数组连乘
func Grow(arr []int) int {
	res := 1
	for _, num := range arr {
		res *= num
	}
	A := strconv.Itoa(res)
	B := fmt.Sprintf("%d", res)
	fmt.Println(A, B, A == B)
	return res
}
func main() {
	fmt.Println(Grow([]int{2, 2, 2, 2, 2, 2}))
}
