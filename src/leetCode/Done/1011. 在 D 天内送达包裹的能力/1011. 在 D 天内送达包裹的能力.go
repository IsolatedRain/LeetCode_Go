// https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days/
package main

import "fmt"

func shipWithinDays(weights []int, D int) int {
	var L, R int
	for _, w := range weights {
		if w > L {
			L = w
		}
		R += w
	}
	for L < R {
		mid := (L + R) >> 1
		if valid(mid, D, weights) {
			R = mid
		} else {
			L++
		}
	}
	return L
}

func valid(shipCapacity, days int, arr []int) bool {
	curSum := 0
	d := 0
	for _, v := range arr {
		curSum += v
		if curSum > shipCapacity {
			curSum = v
			d++
		}
		if d > days {
			return false
		}
	}
	if curSum != 0 {
		d++
	}
	return d <= days
}

func main() {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	D := 5
	fmt.Println(shipWithinDays(weights, D))
}
