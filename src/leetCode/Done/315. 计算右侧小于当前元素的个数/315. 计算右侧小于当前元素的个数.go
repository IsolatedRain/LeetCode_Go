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
	return &BIT{n, make([]int, n)}
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

func countSmaller(nums []int) []int {
	mapSet := map[int]struct{}{}
	for _, v := range nums {
		mapSet[v] = struct{}{}
	}
	arr := []int{}
	for k := range mapSet {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	d := map[int]int{}
	for i := range arr {
		d[arr[i]] = i + 1
	}
	n := len(nums)
	b := newBIT(n + 1)
	res := []int{}
	for i := n - 1; i >= 0; i-- {
		x := d[nums[i]]
		res = append(res, b.querry(x-1))
		b.update(x)
	}
	reverse(res)
	return res
}
func reverse(arr []int) {
	L, R := 0, len(arr)-1
	for L < R {
		arr[L], arr[R] = arr[R], arr[L]
		L++
		R--
	}
}

func main() {
	fmt.Println(countSmaller([]int{5, 2, 6, 1}))
}
