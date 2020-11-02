package main

import "fmt"

func destCity(paths [][]string) string {
	graph := map[string]struct{}{}
	for _, v := range paths {
		graph[v[0]] = struct{}{}
	}
	for _, v := range paths {
		if _, ok := graph[v[1]]; !ok {
			return v[1]
		}
	}
	return ""
}

func main() {
	paths := [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}
	r := destCity(paths)
	fmt.Printf("输入: %v\n输出: %v\n", paths, r)
}
