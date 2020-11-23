package main

import (
	"fmt"
)

func canThreePartsEqualSum(A []int) bool {
	sum := 0
	for _, v := range A {
		sum += v
	}
	if sum%3 != 0 {
		return false
	}
	n := len(A)
	oneThird := sum / 3
	curSum := 0
	k := 0
	i := 0
	for i < n {
		curSum += A[i]
		if curSum == oneThird {
			k++
			curSum = 0
		}
		if k == 2 {
			break
		}
		i++
	}
	return k == 2 && i+1 != n
}

func main() {
	// A := []int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}
	A := []int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1}
	r := canThreePartsEqualSum(A)
	fmt.Printf("%v\n", r)

}
