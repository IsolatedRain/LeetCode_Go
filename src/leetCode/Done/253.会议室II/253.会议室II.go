package main

import (
	"fmt"
	"sort"
)

func minMeetingRooms(intervals [][]int) int {
	nums := []int{}
	vals := map[int]int{}
	for _, i := range intervals {
		vals[i[0]]++
		vals[i[1]]--
	}
	for k := range vals {
		nums = append(nums, k)
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res := 0
	cur := 0
	for _, k := range nums {
		res = max(res, cur)
		cur += vals[k]
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
	intervals := [][]int{{0, 30}, {5, 10}, {15, 20}}
	r := minMeetingRooms(intervals)
	fmt.Println(r)
}
