package main

import (
	"fmt"
	"sort"
)

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	degrees := map[int]int{}
	pre := map[int][]int{}
	for i := 0; i < n; i++ {
		for _, k := range graph[i] {
			pre[k] = append(pre[k], i)
			degrees[i]++
		}
	}
	q := []int{}
	for i := 0; i < n; i++ {
		if degrees[i] == 0 {
			q = append(q, i)
		}
	}
	res := []int{}
	for len(q) > 0 {
		curLen := len(q)
		for _, v := range q {
			res = append(res, v)
			for _, p := range pre[v] {
				degrees[p]--
				if degrees[p] == 0 {
					q = append(q, p)
				}
			}
		}
		q = q[curLen:]
	}
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	return res
}

func main() {
	graph := [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}
	fmt.Println(eventualSafeNodes(graph))
}
