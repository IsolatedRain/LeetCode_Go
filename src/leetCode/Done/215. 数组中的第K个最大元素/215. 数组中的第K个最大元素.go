package main

import (
	"fmt"
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	var quickSort func(arr []int, L, R, tarI int) int
	var randomPartition func(arr []int, L, R int) int
	var partition func(arr []int, L, R int) int

	partition = func(arr []int, L, R int) int {
		pivot := arr[R]
		i := L - 1
		for j := L; j < R; j++ {
			if arr[j] <= pivot {
				i++
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		arr[i+1], arr[R] = arr[R], arr[i+1]
		return i + 1
	}

	randomPartition = func(arr []int, L, R int) int {
		i := rand.Int()%(R-L+1) + L
		arr[i], arr[R] = arr[R], arr[i]
		return partition(arr, L, R)
	}

	quickSort = func(arr []int, L, R, tarI int) int {
		curI := randomPartition(arr, L, R)
		if curI == tarI {
			return arr[curI]
		} else if curI > tarI {
			return quickSort(arr, L, curI-1, tarI)
		}
		return quickSort(arr, curI, R, tarI)
	}
	return quickSort(nums, 0, len(nums)-1, k+1)
}

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	fmt.Println(findKthLargest(nums, k))
}
