package main

import "sort"

// KthLargest 类
type KthLargest struct {
	arr []int
	K   int
}

// Constructor 初始化
func Constructor(k int, nums []int) KthLargest {
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	return KthLargest{nums, k}
}

// Add 插入元素， 返回第 K大
func (k *KthLargest) Add(val int) int {
	idx := sort.Search(len(k.arr), func(i int) bool {
		return k.arr[i] <= val
	})
	tmp := append([]int{}, k.arr[idx:]...)
	k.arr = append(append(k.arr[:idx], val), tmp...)
	v := k.arr[k.K-1]
	k.arr = append(k.arr[:k.K-1], k.arr[k.K:]...)
	return v
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
