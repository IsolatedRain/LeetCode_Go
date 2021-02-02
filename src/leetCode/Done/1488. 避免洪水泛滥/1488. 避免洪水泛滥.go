package main

import (
	"container/heap"
	"fmt"
)

type hp struct {
	days [][]int // {{下一次下雨索引， 当前湖泊编号}}
}

// 按 下一次 最早下雨的索引 排序
func (h *hp) Less(i, j int) bool { return h.days[i][0] < h.days[j][0] }
func (h *hp) Len() int           { return len(h.days) }
func (h *hp) Swap(i, j int) {
	h.days[i], h.days[j] = h.days[j], h.days[i]
	return
}
func (h *hp) Push(v interface{}) { h.days = append(h.days, v.([]int)) }
func (h *hp) Pop() interface{} {
	a := h.days
	v := a[len(a)-1]
	h.days = a[:len(a)-1]
	return v
}
func (h *hp) push(v []int) { heap.Push(h, v) }
func (h *hp) pop() []int   { return heap.Pop(h).([]int) }

func avoidFlood(rains []int) []int {
	h := &hp{}
	heap.Init(h)

	n := len(rains)
	nextRainDay := make([]int, n) // 标记下一次下雨的索引
	mark := map[int]int{}
	for i := range rains {
		if rains[i] != 0 {
			if j, ok := mark[rains[i]]; ok {
				nextRainDay[j] = i
			}
		}
		nextRainDay[i] = n + 1 // 默认置为最后
		mark[rains[i]] = i
	}

	isFull := map[int]bool{} // 标记是否满

	res := make([]int, n) // 预设为 1
	for i := range res {
		res[i] = 1
	}

	for i, v := range rains {
		if rains[i] != 0 {
			res[i] = -1    // 如果下雨
			if isFull[v] { // 如果满了
				return []int{}
			}
			// 按下一次 最早下雨的湖泊 排序
			h.push([]int{nextRainDay[i], rains[i]})
			isFull[v] = true
		} else {
			if len(h.days) > 0 {
				// 弹出下一次 最早下雨的湖泊编号
				dry := h.pop()
				res[i] = dry[1]
				isFull[dry[1]] = false
			}
		}
	}
	return res
}

func main() {
	rains := []int{1, 2, 0, 0, 2, 1}
	fmt.Println(avoidFlood(rains))
}
