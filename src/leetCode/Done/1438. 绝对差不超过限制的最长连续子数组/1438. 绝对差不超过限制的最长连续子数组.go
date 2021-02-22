package main

import "fmt"

func longestSubarray(nums []int, limit int) int {
	L := 0
	res := 1
	maxQ, minQ := []int{}, []int{}
	for R, v := range nums {
		for len(maxQ) > 0 && maxQ[len(maxQ)-1] < v {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, v)
		for len(minQ) > 0 && minQ[len(minQ)-1] > v {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, v)
		for len(maxQ) > 0 && len(minQ) > 0 && maxQ[0]-minQ[0] > limit {
			if maxQ[0] == nums[L] {
				maxQ = maxQ[1:]
			}
			if minQ[0] == nums[L] {
				maxQ = maxQ[1:]
			}
			L++
		}
		if R-L+1 > res {
			res = R - L + 1
		}
	}
	return res
}

func main() {
	nums := []int{10, 1, 2, 4, 7, 2}
	limit := 5
	fmt.Println(longestSubarray(nums, limit))
}
