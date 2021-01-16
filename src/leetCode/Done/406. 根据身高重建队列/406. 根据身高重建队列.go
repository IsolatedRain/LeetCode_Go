package main

import (
	"fmt"
	"sort"
)

func reconstructQueue(people [][]int) (q [][]int) {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0] || (people[i][0] == people[j][0] && people[i][1] < people[j][1])
	})

	insert := func(i int, v []int) {
		q = append(q[:i], append([][]int{v}, q[i:]...)...)
		return
	}

	for i := range people {
		curPos := people[i][1]
		insert(curPos, people[i])
	}
	return
}

func main() {
	// people := [][]int{{6, 0}, {5, 0}, {4, 0}, {3, 2}, {2, 2}, {1, 4}}
	people := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	fmt.Println(reconstructQueue(people))
}
