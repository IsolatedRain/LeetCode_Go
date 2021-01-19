package main

import (
	"fmt"
	"strconv"
)

func findMissingRanges(nums []int, lower int, upper int) []string {
	intervals := [][]int{}
	L := lower - 1
	for _, v := range nums {
		if v <= upper {
			if v-L > 1 {
				intervals = append(intervals, []int{L + 1, v - 1})
			}
		}
		L = v
	}
	if L < upper {
		intervals = append(intervals, []int{L + 1, upper})
	}

	res := []string{}
	for _, inter := range intervals {
		if inter[0] == inter[1] {
			res = append(res, strconv.Itoa(inter[0]))
		} else {
			res = append(res, strconv.Itoa(inter[0])+"->"+strconv.Itoa(inter[1]))
		}
	}
	return res
}

func main() {
	nums := []int{0, 1, 3, 50, 75}
	lower := 0
	upper := 99
	fmt.Println(findMissingRanges(nums, lower, upper))
}
