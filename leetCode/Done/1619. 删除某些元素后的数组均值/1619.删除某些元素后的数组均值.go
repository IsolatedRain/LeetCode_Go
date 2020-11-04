package main

import (
	"fmt"
	"sort"
)

func trimMean(arr []int) float64 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	n := len(arr)
	L, R := n/20, n/20
	sum := 0
	for _, v := range arr[L : n-R] {
		sum += v
	}
	return float64(sum) / float64(n-R-L)
}

func main() {
	// p := lcl.List2GoList("[1,2,2,2,2,2,2,2,2,2,2, 3,2,2,2,2,2,2,2,2]")
	p := []int{9, 7, 8, 7, 7, 8, 4, 4, 6, 8, 8, 7, 6, 8, 8, 9, 2, 6, 0, 0, 1, 10, 8, 6, 3, 3, 5, 1, 10, 9, 0, 7, 10, 0, 10, 4, 1, 10, 6, 9, 3, 6, 0, 0, 2, 7, 0, 6, 7, 2, 9, 7, 7, 3, 0, 1, 6, 1, 10, 3}
	fmt.Printf("输入: %v\n", p)
	r := trimMean(p)
	fmt.Printf("输出: %v\n", r)
}
