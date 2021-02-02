package main

import (
	"fmt"
	"sort"
)

//BIT 树状数组
type BIT struct {
	n    int
	nums []int
}

func newBIT(n int) *BIT {
	return &BIT{n + 1, make([]int, n+1)}
}

func (b *BIT) update(x int) {
	for x < b.n {
		b.nums[x]++
		x += x & -x
	}
}
func (b *BIT) querry(x int) int {
	res := 0
	for x > 0 {
		res += b.nums[x]
		x -= x & -x
	}
	return res
}
func reversePairs(nums []int) int {
	mapSet := map[int]struct{}{}
	for _, v := range nums {
		mapSet[v] = struct{}{}
		mapSet[v*2] = struct{}{}
	}
	arr := []int{}
	for k := range mapSet {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	d := map[int]int{}
	for i, v := range arr {
		d[v] = i + 1
	}
	n := len(nums)
	b := newBIT(len(arr) + 1)
	fmt.Println(d)
	fmt.Println(arr)
	res := 0
	for i := 0; i < n; i++ {
		left := b.querry(d[nums[i]*2])
		right := b.querry(len(arr))
		res += right - left
		b.update(d[nums[i]])
	}
	return res
}

func main() {
	// fmt.Println(reversePairs([]int{2, 4, 3, 5, 1}))
	// fmt.Println(reversePairs([]int{1, 3, 2, 3, 1}))
	nums := []int{1, 3, 2, 3, 1, 2, 3, 4, 5, 6, 2, 3, 4, 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 4, 3, 3, 1}
	fmt.Println(reversePairs(nums))
}
