package main

import (
	"fmt"
	"sort"
)

func highFive(items [][]int) [][]int {
	ids := []int{}
	sort.Slice(items, func(i, j int) bool {
		if items[i][0] == items[j][0] {
			return items[i][1] > items[j][1]
		}
		return items[i][0] < items[j][0]
	})
	mapScores := map[int][]int{}
	res := [][]int{}
	for _, v := range items {
		mapScores[v[0]] = append(mapScores[v[0]], v[1])
		if len(mapScores[v[0]]) == 1 {
			ids = append(ids, v[0])
		}
	}
	for _, id := range ids {
		res = append(res, []int{id, sum(mapScores[id][:5]...) / 5})
	}
	return res
}

func sum(arr ...int) int {
	res := 0
	for _, v := range arr {
		res += v
	}
	return res
}

func main() {
	// items := [][]int{{1, 91}, {1, 92}, {2, 93}, {2, 97}, {1, 60}, {2, 77}, {1, 65}, {1, 87}, {1, 100}, {2, 100}, {2, 76}}
	items := [][]int{{1, 50}, {7, 30}, {1, 100}, {7, 40}, {1, 80}, {7, 70}, {1, 60}, {7, 20}, {1, 90}, {7, 100}}
	r := highFive(items)
	fmt.Println("输出：", r)
}
