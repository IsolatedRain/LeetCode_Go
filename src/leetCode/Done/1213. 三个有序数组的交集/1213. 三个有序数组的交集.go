package main

import "fmt"

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	p1, p2, p3 := 0, 0, 0
	base := min(arr1[p1], arr2[p2], arr3[p3])
	res := []int{}
	for p1 < len(arr1) && p2 < len(arr2) && p3 < len(arr3) {
		if arr1[p1] == base && arr2[p2] == base && arr3[p3] == base {
			res = append(res, base)
			p1++
			p2++
			p3++
		} else {
			if arr1[p1] == base {
				p1++
			}
			if arr2[p2] == base {
				p2++
			}
			if arr3[p3] == base {
				p3++
			}
		}
		if p1 < len(arr1) && p2 < len(arr2) && p3 < len(arr3) {
			base = min(arr1[p1], arr2[p2], arr3[p3])
		}
	}
	return res
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{1, 2, 5, 7, 9}
	arr3 := []int{1, 3, 4, 5, 8}
	fmt.Println(arraysIntersection(arr1, arr2, arr3))
}
func min(x ...int) int {
	r := x[0]
	for _, v := range x[1:] {
		if v < r {
			r = v
		}
	}
	return r
}
