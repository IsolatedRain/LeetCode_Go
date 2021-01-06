package main

import "fmt"

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := map[string][]string{}
	vals := map[string]float64{}
	for i, e := range equations {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
		vals[e[0]+e[1]] = values[i]
		vals[e[1]+e[0]] = 1 / values[i]
	}
	var dfs func(x, y string, val float64)
	get := false
	check := map[string]struct{}{}
	res := []float64{}
	dfs = func(x, y string, val float64) {
		if v, ok := vals[x+y]; ok {
			get = true
			res = append(res, v*val)
			return
		}
		for _, nextNode := range graph[x] {
			if _, ok := check[nextNode+y]; !ok && !get {
				check[nextNode+y] = struct{}{}
				dfs(nextNode, y, val*vals[x+nextNode])
			}
		}
	}
	for _, q := range queries {
		x, y := q[0], q[1]
		_, okX := graph[x]
		_, okY := graph[y]
		if !okX || !okY {
			res = append(res, float64(-1))
			continue
		}
		if x == y {
			res = append(res, float64(1))
			continue
		}
		check = map[string]struct{}{x + y: {}}
		get = false
		dfs(x, y, 1)
		if !get {
			fmt.Println(x, y)
			res = append(res, -1)
		}
	}
	return res
}

func main() {
	// equations := [][]string{{"a", "b"}, {"b", "c"}}
	// values := []float64{2.0, 3.0}
	// queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	equations := [][]string{{"x1", "x2"}, {"x2", "x3"}, {"x3", "x4"}, {"x4", "x5"}}
	values := []float64{3.0, 4.0, 5.0, 6.0}
	queries := [][]string{{"x1", "x5"}, {"x5", "x2"}, {"x2", "x4"}, {"x2", "x2"}, {"x2", "x9"}, {"x9", "x9"}}
	fmt.Println(calcEquation(equations, values, queries))
}
