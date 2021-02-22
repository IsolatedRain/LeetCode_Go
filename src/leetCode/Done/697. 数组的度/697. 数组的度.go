package main

import "fmt"

func findShortestSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count := map[int]int{}
	maxD := 1
	left, right := map[int]int{}, map[int]int{}
	for i, num := range nums {
		count[num]++
		if count[num] > maxD {
			maxD = count[num]
		}
		if _, ok := left[num]; !ok {
			left[num] = i
		}
		right[num] = i
	}

	res := len(nums)
	for _, num := range nums {
		if count[num] == maxD {
			if right[num]-left[num]+1 < res {
				res = right[num] - left[num] + 1
			}
		}
	}
	return res
}

func main() {
	nums := []int{1, 2, 2, 3, 1, 4, 2}
	fmt.Println(findShortestSubArray(nums))
}
