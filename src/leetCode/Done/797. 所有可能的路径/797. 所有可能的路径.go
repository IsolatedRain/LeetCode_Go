package main

import "fmt"

func allPathsSourceTarget(graph [][]int) [][]int {
	res := [][]int{}
	tar := len(graph) - 1
	var dfs func(node int, path []int)
	dfs = func(node int, path []int) {
		if node == tar {
			res = append(res, append([]int{}, path...))
			return
		}
		for _, nextNde := range graph[node] {
			dfs(nextNde, append(path, nextNde))
		}
	}
	dfs(0, []int{0})
	return res
}

func main() {
	graph := [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}
	fmt.Println(allPathsSourceTarget(graph))

}
