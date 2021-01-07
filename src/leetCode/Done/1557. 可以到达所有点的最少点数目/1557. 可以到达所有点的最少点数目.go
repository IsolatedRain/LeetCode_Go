package main

import "fmt"

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	indegree := make([]int, n)
	for _, v := range edges {
		indegree[v[1]]++
	}
	res := []int{}
	for i := range indegree {
		if indegree[i] == 0 {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	n := 5
	edges := [][]int{{0, 1}, {2, 1}, {3, 1}, {1, 4}, {2, 4}}
	fmt.Println(findSmallestSetOfVertices(n, edges))
}
