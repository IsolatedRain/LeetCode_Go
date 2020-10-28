package main

import "fmt"

// 小的有几个, 就向右偏移 几位.
func findKthPositive(arr []int, k int) int {
	for _, v := range arr {
		if v > k {
			return k
		}
		k++
	}
	return k
}

func main() {
	arr := []int{2, 3, 4, 7, 11}
	k := 5
	fmt.Printf("%v\n", findKthPositive(arr, k))
}
