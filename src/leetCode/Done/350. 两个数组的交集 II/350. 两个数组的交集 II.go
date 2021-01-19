package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	var mergeSort func([]int) []int
	mergeSort = func(arr []int) []int {
		n := len(arr)
		if n < 2 {
			return arr
		}
		left := mergeSort(arr[:n/2])
		right := mergeSort(arr[n/2:])
		L, R := 0, 0
		r := []int{}
		for L < len(left) && R < len(right) {
			if left[L] > right[R] {
				r = append(r, right[R])
				R++
			} else {
				r = append(r, left[L])
				L++
			}
		}
		if L == len(left) {
			r = append(r, right[R:]...)
		} else {
			r = append(r, left[L:]...)
		}
		return r
	}
	nums1 = mergeSort(nums1)
	nums2 = mergeSort(nums2)
	p1, p2 := 0, 0
	res := []int{}
	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] == nums2[p2] {
			res = append(res, nums1[p1])
			p1++
			p2++
		} else {
			if nums1[p1] < nums2[p2] {
				p1++
			} else {
				p2++
			}
		}
	}
	return res
}

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersect(nums1, nums2))
}
