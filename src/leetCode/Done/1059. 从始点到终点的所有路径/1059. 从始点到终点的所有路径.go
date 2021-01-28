package main

import "fmt"

func leadsToDestination(n int, edges [][]int, source int, destination int) bool {
	if len(edges) == 0 {
		return source == destination
	}
	graph := map[int][]int{}
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
	}
	visited := map[int]bool{}
	var dfs func(node int) bool
	dfs = func(node int) bool {
		if _, ok := graph[node]; !ok {
			if node != destination {
				return false
			}
		}
		for _, nextNode := range graph[node] {
			if !visited[nextNode] {
				visited[nextNode] = true
				if !dfs(nextNode) {
					return false
				}
				visited[nextNode] = false
			} else {
				return false
			}
		}
		return true
	}
	return dfs(source)
}

func main() {
	fmt.Println(leadsToDestination(3, [][]int{{0, 1}, {0, 2}}, 0, 2))
}
