package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 建图 和 前置课程数,即 入度数
	graph := make(map[int][]int, 0)
	inDegree := make([]int, numCourses, numCourses)
	for _, v := range prerequisites {
		graph[v[1]] = append(graph[v[1]], v[0])
		inDegree[v[0]]++
	}

	// 从前置课程为0的课程开始
	q := make([]int, 0, numCourses)
	for i, v := range inDegree {
		if v == 0 {
			q = append(q, i)
		}
	}

	count := len(q)
	for len(q) > 0 {
		curCourse := q[0]
		q = q[1:]
		for _, nextCourse := range graph[curCourse] {
			inDegree[nextCourse]--
			if inDegree[nextCourse] == 0 {
				q = append(q, nextCourse)
				count++
			}
		}
	}

	return count == numCourses
}

func main() {
	// numCourses := 3
	// prerequisites := [][]int{{1, 0}, {1, 2}, {0, 1}}
	numCourses := 2
	prerequisites := [][]int{{1, 0}}
	fmt.Printf("输入: %v,  %#v\n", numCourses, prerequisites)
	r := canFinish(numCourses, prerequisites)
	fmt.Printf("输出: %v\n", r)
}
