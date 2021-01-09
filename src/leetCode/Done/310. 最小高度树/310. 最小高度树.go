package main

import "fmt"

func findMinHeightTrees(n int, edges [][]int) []int {
	graph := map[int][]int{}
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	startNode := -1
	for k := range edges {
		if len(graph[k]) == 1 {
			startNode = k
			break
		}
	}

	maxLen := 1
	visited := map[int]bool{}
	visited[startNode] = true
	path := []int{}
	c := 0
	var findLongestPath func(curNode int, curLen int, curPath []int)
	findLongestPath = func(curNode int, curLen int, curPath []int) {
		if len(graph[curNode]) == 1 && curLen > maxLen {
			maxLen, startNode, path = curLen, curNode, curPath
			return
		}
		for _, nextNode := range graph[curNode] {
			if !visited[nextNode] {
				c++
				visited[nextNode] = true
				findLongestPath(nextNode, curLen+1, append(curPath, nextNode))
				visited[nextNode] = false
			}
		}
	}
	findLongestPath(startNode, 1, []int{startNode})
	visited = map[int]bool{}
	visited[startNode] = true
	maxLen = 1
	findLongestPath(startNode, 1, []int{startNode})
	if maxLen&1 == 1 {
		return []int{path[maxLen/2]}
	}
	return []int{path[maxLen/2], path[maxLen/2-1]}
}

func main() {
	n := 6
	edges := [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}
	fmt.Println(findMinHeightTrees(n, edges))
}
