package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	toStr := []string{}
	for _, num := range nums {
		toStr = append(toStr, strconv.Itoa(num))
	}
	sort.Slice(toStr, func(i, j int) bool {
		return toStr[i]+toStr[j] > toStr[j]+toStr[i]
	})
	res := strings.Join(toStr, "")
	if res[0] == '0' {
		return "0"
	}
	return res
}

func main() {
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
}

// func largestNumber(nums []int) string {
// 	sort.Slice(nums, func(i, j int) bool {
// 		a, b := nums[i], nums[j]
// 		x, y := 10, 10
// 		for x <= a {
// 			x *= 10
// 		}
// 		for y <= b {
// 			y *= 10
// 		}
// 		return a*y+b > b*x+a
// 	})
// 	toStr := []string{}
// 	for _, num := range nums {
// 		toStr = append(toStr, strconv.Itoa(num))
// 	}
// 	if toStr[0][0] == '0' {
// 		return "0"
// 	}
// 	return strings.Join(toStr, "")
// }
