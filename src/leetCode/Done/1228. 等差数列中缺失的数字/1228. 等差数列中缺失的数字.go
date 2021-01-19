package main

import "fmt"

func missingNumber(arr []int) int {
	n := len(arr)
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	return (arr[0]+arr[n-1])*(n+1)/2 - sum
}

func main() {
	arr := []int{5, 7, 11, 13}
	fmt.Println(missingNumber(arr))
}
