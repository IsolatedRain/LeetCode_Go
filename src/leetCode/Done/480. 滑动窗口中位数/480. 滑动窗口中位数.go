package main

import (
	"fmt"
	"sort"
)

func medianSlidingWindow(nums []int, k int) []float64 {
	del := func(arr []int, v int) []int {
		idx := sort.Search(len(arr), func(i int) bool {
			return arr[i] >= v
		})
		arr = append(append([]int{}, arr[:idx]...), arr[idx+1:]...)
		return arr
	}

	insert := func(arr []int, v int) []int {
		idx := sort.Search(len(arr), func(i int) bool {
			return arr[i] >= v
		})
		arr = append(arr[:idx], append([]int{v}, arr[idx:]...)...)
		return arr
	}
	getMid := func(arr []int) float64 {
		if k&1 == 1 {
			return float64(arr[k/2])
		}
		return float64((arr[k/2] + arr[(k-1)/2])) / 2.0
	}
	arr := append([]int{}, nums[:k]...)
	sort.Ints(arr)
	res := []float64{}
	res = append(res, getMid(arr))
	n := len(nums)
	L := 0
	for R := k; R < n; R++ {
		arr = del(arr, nums[L])
		arr = insert(arr, nums[R])
		L++
		res = append(res, getMid(arr))
	}
	return res
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(medianSlidingWindow(nums, k))
}
