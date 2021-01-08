package main

import "fmt"

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	graph := map[int][]int{}
	p := map[[2]int]float64{}
	for i, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
		p[[2]int{e[0], e[1]}] = succProb[i]
		p[[2]int{e[1], e[0]}] = succProb[i]
	}
	dp := make([]float64, n)
	visited := make([]bool, n)
	q := [][]float64{{float64(start), 1}}
	for len(q) > 0 {
		curLen := len(q)
		for i := 0; i < curLen; i++ {
			curNode, curP := int(q[i][0]), q[i][1]
			if curP > dp[curNode] {
				dp[curNode] = curP
			}
			visited[curNode] = true
			for _, nextNode := range graph[curNode] {
				if p[[2]int{curNode, nextNode}]*curP > dp[nextNode] {
					q = append(q, []float64{float64(nextNode), p[[2]int{curNode, nextNode}] * curP})
				}
			}
		}
		q = q[curLen:]
	}
	return dp[end]
}

func main() {
	n := 3
	edges := [][]int{{0, 1}, {1, 2}, {0, 2}}
	succProb := []float64{0.5, 0.5, 0.2}
	start := 0
	end := 2
	fmt.Println(maxProbability(n, edges, succProb, start, end))
}
