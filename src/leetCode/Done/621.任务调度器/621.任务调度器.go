package main

import "fmt"

// A B -
// A B -
// A B
func leastInterval(tasks []byte, n int) int {
	cnt := map[byte]int{}
	maxCount := 0
	for _, b := range tasks {
		cnt[b]++
		if cnt[b] > maxCount {
			maxCount = cnt[b]
		}
	}
	minCost := (maxCount - 1) * (n + 1)
	for _, v := range cnt {
		if v == maxCount {
			minCost++
		}
	}
	if minCost < len(tasks) {
		return len(tasks)
	}
	return minCost
}

func main() {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n := 2
	r := leastInterval(tasks, n)
	fmt.Println(r)
}
