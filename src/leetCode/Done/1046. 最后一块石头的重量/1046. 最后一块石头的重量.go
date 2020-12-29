package main

import (
	"fmt"
	"sort"
)

func lastStoneWeight(stones []int) int {
	sort.Ints(stones)
	for len(stones) > 1 {
		a := stones[len(stones)-1]
		b := stones[len(stones)-2]
		stones = stones[:len(stones)-2]
		newStone := a - b
		if newStone != 0 {
			stones = append(stones, newStone)
			sort.Ints(stones)
		}
	}
	if len(stones) > 0 {
		return stones[0]
	}
	return 0
}

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	r := lastStoneWeight(stones)
	fmt.Println(r)
}

// type hp struct {
// 	sort.IntSlice
// }

// func (h *hp) Less(i, j int) bool {
// 	return h.IntSlice[i] > h.IntSlice[j]
// }

// func (h *hp) Push(v interface{}) {
// 	h.IntSlice = append(h.IntSlice, v.(int))
// }

// func (h *hp) Pop() interface{} {
// 	a := h.IntSlice
// 	v := a[len(a)-1]
// 	h.IntSlice = a[:len(a)-1]
// 	return v
// }

// func (h *hp) push(v int) {
// 	heap.Push(h, v)
// }

// func (h *hp) pop() int {
// 	return heap.Pop(h).(int)

// }

// func lastStoneWeight(stones []int) int {
// 	q := &hp{stones}
// 	heap.Init(q)
// 	for len(q.IntSlice) > 1 {
// 		a := q.pop()
// 		b := q.pop()
// 		if a-b > 0 {
// 			q.push(a - b)
// 		}
// 	}
// 	if len(q.IntSlice) > 0 {
// 		return q.IntSlice[0]
// 	}
// 	return 0
// }
