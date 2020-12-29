package main

import "fmt"

func numOfSubarrays(arr []int) int {
	odd := 0
	even := 1
	curSum := 0
	res := 0
	n := len(arr)
	mod := 1000000007
	for i := 0; i < n; i++ {
		curSum += arr[i]
		if curSum&1 == 1 {
			res += even
			odd++
		} else {
			res += odd
			even++
		}
		res %= mod
	}
	return res
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	// arr := []int{1, 3, 5}
	r := numOfSubarrays(arr)
	fmt.Println(r)
}
