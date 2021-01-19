package main

import "fmt"

func transformArray(arr []int) []int {
	n := len(arr)
	transed := true
	for transed {
		transed = false
		for i := 1; i < n-1; i++ {
			if arr[i] < arr[i-1] && arr[i] < arr[i+1] {
				transed = true
				arr[i]++
			} else if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
				transed = true
				arr[i]--
			}
		}
	}
	return arr
}

func main() {
	arr := []int{1, 6, 3, 4, 3, 5}
	fmt.Println(transformArray(arr))
}
