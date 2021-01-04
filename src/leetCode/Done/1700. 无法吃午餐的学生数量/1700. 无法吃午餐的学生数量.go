package main

import "fmt"

func countStudents(students []int, sandwiches []int) int {
	s := [2]int{}
	for _, v := range students {
		s[v]++
	}
	for i := range sandwiches {
		if s[sandwiches[i]] == 0 {
			return len(sandwiches) - i
		}
		s[sandwiches[i]]--
	}
	return 0
}

func main() {
	students := []int{1, 1, 1, 0, 0, 1}
	sandwiches := []int{1, 0, 0, 0, 1, 1}
	fmt.Println(countStudents(students, sandwiches))
}
