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

func maxEvents(events [][]int) int {
	maxDay := 0
	for _, e := range events {
		if e[1] > maxDay {
			maxDay = e[1]
		}
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0] || (events[i][0] == events[j][0] && events[i][1] < events[j][1])
	})
	today := &hp{}
	heap.Init(today)
	res := 0
	for i := 1; i < maxDay+1; i++ {
		for len(events) > 0 && events[0][0] <= i {
			today.push(events[0][1])
			events = events[1:]
		}
		for len(today.IntSlice) > 0 {
			endDay := today.pop()
			if endDay >= i {
				res++
				break
			}
		}
		if len(events) == 0 && len(today.IntSlice) == 0 {
			break
		}
	}
	return res
}

func main() {
	// events := [][]int{{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6}, {1, 7}}
	events := [][]int{{1, 4}, {4, 4}, {2, 2}, {3, 4}, {1, 1}}
	// events := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 2}}
	// events := [][]int{{1, 2}, {2, 3}, {3, 4}}
	fmt.Println(maxEvents(events))
}
