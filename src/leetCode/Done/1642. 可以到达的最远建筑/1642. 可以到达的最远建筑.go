package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type hp struct{ sort.IntSlice }

func (h *hp) Less(i, j int) bool { return h.IntSlice[i] < h.IntSlice[j] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func (h *hp) push(v int) { heap.Push(h, v) }
func (h *hp) pop() int   { return heap.Pop(h).(int) }

func furthestBuilding(heights []int, bricks int, ladders int) int {
	q := &hp{}
	heap.Init(q)
	n := len(heights)
	for i := 1; i < n; i++ {
		if heights[i] > heights[i-1] {
			q.push(heights[i] - heights[i-1])
			if len(q.IntSlice) > ladders {
				bricks -= q.pop()
				if bricks < 0 {
					return i - 1
				}
			}
		}
	}
	return n - 1
}

func main() {
	// heights := []int{4, 2, 7, 6, 9, 14, 12}
	// bricks := 5
	// ladders := 1
	heights := []int{4, 12, 2, 7, 3, 18, 20, 3, 19}
	bricks := 10
	ladders := 2
	r := furthestBuilding(heights, bricks, ladders)
	fmt.Println(r)
}
