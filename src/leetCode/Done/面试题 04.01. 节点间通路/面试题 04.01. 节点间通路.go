package main

import "fmt"

func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	d := map[int][]int{}
	for _, v := range graph {
		d[v[0]] = append(d[v[0]], v[1])
	}
	visited := map[int]bool{}
	visited[start] = true
	q := []int{start}
	for len(q) > 0 {
		curNode := q[0]
		if curNode == target {
			return true
		}
		q = q[1:]
		for _, nextNode := range d[curNode] {
			if !visited[nextNode] {
				visited[nextNode] = true
				q = append(q, nextNode)
			}
		}
	}
	return false
}

func main() {
	n := 5
	graph := [][]int{{0, 1}, {0, 2}, {0, 4}, {0, 4}, {0, 1}, {1, 3}, {1, 4}, {1, 3}, {2, 3}, {3, 4}}
	start := 0
	target := 4
	fmt.Println(findWhetherExistsPath(n, graph, start, target))
}
