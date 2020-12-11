package main

import "fmt"

func canReach(arr []int, start int) bool {
	visited := map[int]struct{}{start: {}}
	q := []int{start}
	n := len(arr)
	for len(q) > 0 {
		if arr[q[0]] == 0 {
			return true
		}
		left := q[0] - arr[q[0]]
		right := q[0] + arr[q[0]]
		if _, ok := visited[left]; !ok && left >= 0 {
			visited[left] = struct{}{}
			q = append(q, left)
		}
		if _, ok := visited[right]; !ok && right < n {
			visited[right] = struct{}{}
			q = append(q, right)
		}
		q = q[1:]
	}
	return false
}

func main() {
	arr := []int{4, 2, 3, 0, 3, 1, 2}
	start := 5
	r := canReach(arr, start)
	fmt.Println(r)
}
