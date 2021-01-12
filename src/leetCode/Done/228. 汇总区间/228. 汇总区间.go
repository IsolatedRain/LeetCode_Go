package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	n := len(nums)
	if n == 0 {
		return []string{}
	}
	if n == 1 {
		return []string{strconv.Itoa(nums[0])}
	}

	nums = append(nums, nums[n-1]+2)
	L, R := 0, 1
	intervals := [][]int{}
	for R < n+1 {
		if nums[R]-nums[R-1] > 1 {
			intervals = append(intervals, []int{nums[L], nums[R-1]})
			L = R
		}
		R++
	}

	res := []string{}
	for _, v := range intervals {
		if v[0] == v[1] {
			res = append(res, strconv.Itoa(v[0]))
		} else {
			res = append(res, strconv.Itoa(v[0])+"->"+strconv.Itoa(v[1]))
		}
	}

	return res
}

func main() {
	nums := []int{0, 2, 3, 4, 6, 8, 9}
	fmt.Println(summaryRanges(nums))
}
