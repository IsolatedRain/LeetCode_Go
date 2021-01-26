package main

import "fmt"

func maxNumEdgesToRemove(n int, edges [][]int) int {
	graph := map[int][][]int{}
	res := 0

	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], []int{e[1], e[2]})
	}

	p := make([]int, n+1)
	for i := range p {
		p[i] = i
	}

	var getRoot func(x int) int
	getRoot = func(x int) int {
		if p[x] != p[p[x]] {
			p[x] = getRoot(p[x])
		}
		return p[x]
	}

	var countEdges func(k int)
	countEdges = func(k int) {
		for _, v := range graph[k] {
			xRoot, yRoot := getRoot(v[0]), getRoot(v[1])
			if xRoot == yRoot {
				res++
				continue
			}
			p[xRoot] = yRoot
		}
	}

	var isValid func() bool
	isValid = func() bool {
		cur := 0
		for i := range p {
			if p[i] == i {
				cur++
			}
		}
		return cur == 2
	}

	countEdges(3)
	p1 := append([]int{}, p...)

	countEdges(2)
	if !isValid() {
		return -1
	}

	p = p1
	countEdges(1)
	if !isValid() {
		return -1
	}
	return res
}

func main() {
	fmt.Println(maxNumEdgesToRemove(4,
		[][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 3}, {1, 2, 4}, {1, 1, 2}, {2, 3, 4}}))
}
