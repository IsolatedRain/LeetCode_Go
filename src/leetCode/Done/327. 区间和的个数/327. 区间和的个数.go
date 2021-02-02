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

func countRangeSum(nums []int, lower int, upper int) int {
	preSum := []int{0}
	for i := range nums {
		preSum = append(preSum, preSum[len(preSum)-1]+nums[i])
	}
	mapSet := map[int]struct{}{}
	for _, v := range preSum {
		mapSet[v] = struct{}{}
		mapSet[v-lower] = struct{}{}
		mapSet[v-upper] = struct{}{}
	}
	arr := []int{}
	for k := range mapSet {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	d := map[int]int{}
	for i, v := range arr {
		d[v] = i
	}
	res := 0
	b := newBIT(len(d))
	for _, v := range preSum {
		left := d[v-upper]
		right := d[v-lower]
		res += b.querry(right+1) - b.querry(left)
		b.update(d[v] + 1)
		fmt.Println(b.nums,d[v])
	}
	return res
}

func main() {
	fmt.Println(countRangeSum([]int{-2, 5, -1}, -2, 2))
}
