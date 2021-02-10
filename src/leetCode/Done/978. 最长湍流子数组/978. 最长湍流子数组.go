package main

import "fmt"

func maxTurbulenceSize(arr []int) int {
	arr = append(arr, arr[len(arr)-1])
	res, L, lastStatus, curStatus, n := 1, 0, 3, 0, len(arr)
	for i := 1; i < n; i++ {
		if arr[i] > arr[i-1] {
			curStatus = 1
		} else if arr[i] < arr[i-1] {
			curStatus = -1
		} else {
			curStatus = 0
		}
		if curStatus == 0 {
			if lastStatus != 0 {
				res = max(res, i-L)
				L = i
			}
		} else if curStatus == lastStatus && lastStatus != 0 {
			res = max(res, i-L)
			L = i - 1
		}
		lastStatus = curStatus
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
	arr := []int{9, 4, 2, 10, 7, 8, 8, 1, 9}
	// arr := []int{0, 1, 1, 0, 1, 0, 1, 1, 0, 0}
	// arr := []int{1, 2, 3, 1}
	fmt.Println(maxTurbulenceSize(arr))
}
