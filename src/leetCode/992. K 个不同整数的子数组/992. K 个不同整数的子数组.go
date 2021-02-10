package main

import (
	"fmt"
)

func subarraysWithKDistinct(A []int, K int) int {
	n, count1, count2, R1, R2 := len(A), map[int]int{}, map[int]int{}, 0, 0
	res := 0
	getKth := func(count map[int]int, k int, i int) int {
		for i < n && len(count) <= k {
			count[A[i]]++
			i++
		}
		if len(count) < k {
			return -1
		} else if len(count) > k {
			i--
			delete(count, A[i])
		}
		return i
	}
	for _, v := range A {
		maxR := getKth(count1, K, R1)
		minR := getKth(count2, K-1, R2)
		if maxR == -1 || minR == -1 {
			return res
		}
		res += maxR - minR
		fmt.Println(count1)
		fmt.Println(count2)
		fmt.Println(maxR, minR, res)
		R1 = maxR
		R2 = minR
		count1[v]--
		if count1[v] == 0 {
			delete(count1, v)
		}
		if len(count2) > 0 {
			count2[v]--
			if count2[v] == 0 {
				delete(count2, v)
			}
		} else {
			R2 ++
		}
	}
	return res
}

func main() {
	// A := []int{1, 2, 1, 2, 3, 1, 2}
	A := []int{2, 1, 1, 1, 2}
	// A := []int{1, 2, 1, 2, 3}
	K := 1
	fmt.Println(subarraysWithKDistinct(A, K))
}
