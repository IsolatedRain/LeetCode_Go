package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make(map[int][]int)
	inDegree := make([]int, numCourses, numCourses)
	for _, v := range prerequisites {
		graph[v[1]] = append(graph[v[1]], v[0])
		inDegree[v[0]]++
	}

	q := make([]int, 0, numCourses)
	for i, v := range inDegree {
		if v == 0 {
			q = append(q, i)
		}
	}

	res := q
	for len(q) > 0 {
		curCourse := q[0]
		q = q[1:]
		for _, nextCourse := range graph[curCourse] {
			inDegree[nextCourse]--
			if inDegree[nextCourse] == 0 {
				res = append(res, nextCourse)
				q = append(q, nextCourse)
			}
		}
	}

	if len(res) != numCourses {
		return []int{}
	}
	return res
}

func main() {
	numCourses := 4
	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	fmt.Printf("输入: %v    %v\n", numCourses, prerequisites)
	r := findOrder(numCourses, prerequisites)
	fmt.Printf("输出: %v\n", r)
}
