package main

import "fmt"

func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	groupItems := map[int][]int{}
	groupID := m
	for i, v := range group {
		if v == -1 {
			group[i] = groupID // 分配组项目一个 小组ID
			groupID++
		}
		groupItems[group[i]] = append(groupItems[group[i]], i)
	}

	// 组间图 和 小组入度
	groupGraph := map[int][]int{}
	groupIndegrees := map[int]int{}
	// 组内图 和 组内项目 入度
	itemsGraph := map[int][]int{}
	itemIndegrees := map[int]int{}

	for item, pres := range beforeItems {
		for _, pre := range pres {
			if group[item] == group[pre] { // 同一小组，则分配组内图与入度
				itemsGraph[pre] = append(itemsGraph[pre], item)
				itemIndegrees[item]++
			} else { // 组间图 与 入度
				groupGraph[group[pre]] = append(groupGraph[group[pre]], group[item])
				groupIndegrees[group[item]]++
			}
		}
	}

	// 拓扑排序
	var topoSort func(graph map[int][]int, degrees map[int]int, items []int) []int
	topoSort = func(graph map[int][]int, degrees map[int]int, items []int) []int {
		q := []int{}
		for _, i := range items {
			if degrees[i] == 0 {
				q = append(q, i)
			}
		}
		res := []int{}
		for len(q) > 0 {
			cur := q[0]
			res = append(res, cur)
			q = q[1:]
			for _, v := range graph[cur] {
				degrees[v]--
				if degrees[v] == 0 {
					q = append(q, v)
				}
			}
		}
		return res
	}

	// 排序小组
	items := make([]int, groupID+1)
	for i := range items {
		items[i] = i
	}
	orderedGroups := topoSort(groupGraph, groupIndegrees, items)
	if len(orderedGroups) < len(items) {
		return []int{}
	}

	// 排序组内
	ans := []int{}
	for _, v := range orderedGroups {
		items = groupItems[v]
		orderedItems := topoSort(itemsGraph, itemIndegrees, items)
		if len(orderedItems) < len(items) {
			return []int{}
		}
		ans = append(ans, orderedItems...)
	}

	return ans
}

func main() {
	n := 8
	m := 2
	group := []int{-1, -1, 1, 0, 0, 1, 0, -1}
	beforeItems := [][]int{{}, {6}, {5}, {6}, {3, 6}, {}, {}, {}}
	fmt.Println(sortItems(n, m, group, beforeItems))
}
