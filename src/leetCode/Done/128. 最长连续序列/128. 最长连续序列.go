package main

import "fmt"

func longestConsecutive(nums []int) int {
	d := map[int]int{}
	maxLen := 0
	for _, v := range nums {
		if _, ok := d[v]; ok {
			continue
		}
		left := 0
		if leftLen, ok := d[v-1]; ok {
			left = leftLen
		}
		right := 0
		if rightLen, ok := d[v+1]; ok {
			right = rightLen
		}
		curLen := left + 1 + right
		if curLen > maxLen {
			maxLen = curLen
		}
		d[v] = curLen
		d[v+right] = curLen
		d[v-left] = curLen
	}
	return maxLen
}

func main() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(nums))
}
