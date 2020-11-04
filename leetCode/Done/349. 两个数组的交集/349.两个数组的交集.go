package main

import (
	"fmt"
)

func intersection(nums1 []int, nums2 []int) []int {
	set1 := map[int]bool{}
	for _, v := range nums1 {
		set1[v] = false
	}
	for _, v := range nums2 {
		if _, ok := set1[v]; ok {
			set1[v] = true
		}
	}
	res := []int{}
	for k, v := range set1 {
		if v == true {
			res = append(res, k)
		}
	}
	return res
}

// func intersection(nums1 []int, nums2 []int) []int {
// 	sort.Ints(nums1)
// 	sort.Ints(nums2)
// 	var res []int
// 	idx1, idx2 := 0, 0
// 	for idx1 < len(nums1) && idx2 < len(nums2) {
// 		x, y := nums1[idx1], nums2[idx2]
// 		if x == y {
// 			if res == nil || x > res[len(res)-1] {
// 				res = append(res, x)
// 			}
// 			idx1++
// 			idx2++
// 		} else if x < y {
// 			idx1++
// 		} else {
// 			idx2++
// 		}
// 	}

// 	return res
// }

func main() {
	p1 := []int{4, 9, 5}
	p2 := []int{9, 4, 9, 8, 4}
	fmt.Printf("输入: %v %v\n", p1, p2)
	r := intersection(p1, p2)
	fmt.Printf("输出: %v\n", r)
}
