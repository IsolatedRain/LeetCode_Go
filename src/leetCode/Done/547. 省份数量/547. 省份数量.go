package main

import "fmt"

//并查集
func findCircleNum(isConnected [][]int) int {
	row := len(isConnected)
	p := make([]int, row)
	for i := range p {
		p[i] = i
	}

	var getRoot func(int) int
	getRoot = func(x int) int {
		if p[x] != x {
			p[x] = getRoot(p[x])
		}
		return p[x]
	}

	var union func(x, y int)
	union = func(x, y int) {
		xRoot := getRoot(x)
		yRoot := getRoot(y)
		if xRoot > yRoot {
			p[xRoot] = yRoot
		} else {
			p[yRoot] = xRoot
		}
	}
	for i, r := range isConnected {
		for j := i + 1; j < row; j++ {
			if r[j] == 1 {
				union(i, j)
			}
		}
	}

	res := 0
	for i, v := range p {
		if i == v {
			res++
		}
	}
	return res
}

func main() {
	// isConnected := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	isConnected := [][]int{
		{1, 0, 0, 1},
		{0, 1, 1, 0},
		{0, 1, 1, 1},
		{1, 0, 1, 1}}
	fmt.Println(findCircleNum(isConnected))
}

// DFS
// func findCircleNum(isConnected [][]int) int {
// 	row := len(isConnected)
// 	cityVisited := make([]bool, row)

// 	var dfs func(curCity int)
// 	dfs = func(curCity int) {
// 		cityVisited[curCity] = true
// 		for city, v := range isConnected[curCity] {
// 			if v == 1 && !cityVisited[city] {
// 				dfs(city)
// 			}
// 		}
// 	}

// 	res := 0
// 	for curCity, v := range cityVisited {
// 		if !v {
// 			res++
// 			dfs(curCity)
// 		}
// 	}
// 	return res
// }
