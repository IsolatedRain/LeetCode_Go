package main

import "fmt"

func fairCandySwap(A []int, B []int) []int {
	aSum, bSum, bCandy := 0, 0, map[int]bool{}
	for i := range A {
		aSum += A[i]
	}
	for i := range B {
		bSum += B[i]
		bCandy[B[i]] = true
	}
	diff := aSum - bSum
	diff /= 2
	for i := range A {
		if _, ok := bCandy[A[i]-diff]; ok {
			return []int{A[i], A[i] - diff}
		}
	}
	return []int{}
}

func main() {
	A := []int{2, 2}
	B := []int{1, 1}
	fmt.Println(fairCandySwap(A, B))
}
