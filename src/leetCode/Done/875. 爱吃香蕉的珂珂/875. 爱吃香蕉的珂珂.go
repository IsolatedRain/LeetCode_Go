// https://leetcode-cn.com/problems/koko-eating-bananas/
package main

import (
	"fmt"
)

func minEatingSpeed(piles []int, h int) int {
	R, sum := 0, 0
	for _, p := range piles {
		if p > R {
			R = p
		}
		sum += p
	}
	L := sum / h
	for L < R {
		mid := (L + R) >> 1
		if eat(piles, h, mid) {
			R = mid
		} else {
			L = mid + 1
		}
	}
	return L
}

func eat(p []int, h, e int) bool {
	curH := 0
	for _, v := range p {
		a, b := v/e, v%e
		if b != 0 {
			a++
		}
		curH += a
	}
	return curH <= h
}

func main() {
	// piles := []int{30, 11, 23, 4, 20}
	// H := 6
	piles := []int{3, 6, 7, 11}
	H := 8
	fmt.Println(minEatingSpeed(piles, H))
}
