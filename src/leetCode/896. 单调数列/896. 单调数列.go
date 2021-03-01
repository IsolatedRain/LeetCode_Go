package main

import (
	"fmt"
)

func isMonotonic(A []int) bool {
	inc, dec := true, true
	for i := 1; i < len(A); i++ {
		if A[i] < A[i-1] {
			inc = false
		}
		if A[i] > A[i-1] {
			dec = false
		}
	}
	return inc || dec
}

func main() {
	fmt.Print([]int{1, 2, 2, 3})
}
