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
func (b *BIT) query(x int) int {
	res := 0
	for x > 0 {
		res += b.nums[x]
		x -= x & -x
	}
	return res
}

func createSortedArray(instructions []int) int {
	mapSet := map[int]struct{}{}
	for _, v := range instructions {
		mapSet[v] = struct{}{}
	}
	allNums := []int{}
	for k := range mapSet {
		allNums = append(allNums, k)
	}
	sort.Ints(allNums)
	d := map[int]int{}
	for i, v := range allNums {
		d[v] = i + 1
	}
	b := newBIT(len(allNums) + 1)
	n := len(instructions)
	res := 0
	mod := 1000000007
	for i := 0; i < n; i++ {
		cur := b.query(d[instructions[i]])
		less := b.query(d[instructions[i]] - 1)
		bigger := i - cur
		res += min(less, bigger) % mod
		b.update(d[instructions[i]])
	}
	return res
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	instructions := []int{1, 2, 3, 6, 5, 4}
	fmt.Println(createSortedArray(instructions))
}
