// https://www.codewars.com/kata/5a6b24d4e626c59d5b000066/train/go
package main

import "fmt"

// SquareSumsRow 和是平方 的路径
func SquareSumsRow(n int) []int {
	sqr := []int{}
	for i := 2; i < n+1; i++ {
		v := i * i
		if v > n*2 {
			break
		}
		sqr = append(sqr, v)
	}
	graph := map[int][]int{}
	for i := 1; i < n+1; i++ {
		for _, v := range sqr {
			if v-i > n {
				break
			}
			if v-i != i && v-i > 0 {
				graph[i] = append(graph[i], v-i)
			}
		}
	}
	visited := map[int]bool{}
	path := make([]int, n)
	var dfs func(i, node int) bool
	dfs = func(i, node int) bool {
		if i == n {
			return true
		}
		for _, nextNode := range graph[node] {
			if !visited[nextNode] {
				visited[nextNode] = true
				path[i] = nextNode
				if dfs(i+1, nextNode) {
					return true
				}
				path[i] = 0
				visited[nextNode] = false
			}
		}
		return false
	}
	exisit := false
	for i := 1; i < n+1; i++ {
		if dfs(0, i) {
			exisit = true
			break
		}
	}
	if !exisit {
		return nil
	}
	return path
}

func main() {
	fmt.Println(SquareSumsRow(15), []int{9, 7, 2, 14, 11, 5, 4, 12, 13, 3, 6, 10, 15, 1, 8})
}
