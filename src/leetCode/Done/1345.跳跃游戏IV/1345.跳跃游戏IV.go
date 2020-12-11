package main

import "fmt"

type queue []int

func (q *queue) append(val int) {
	*q = append(*q, val)
}
func (q *queue) popleft() int {
	r := (*q)[0]
	*q = (*q)[1:]
	return r
}
func newQueue() queue {
	return queue{}
}

func minJumps(arr []int) int {
	n := len(arr)
	if n <= 1 {
		return 0
	}
	idxMap := map[int][]int{}
	nums := []int{arr[0]}
	for i := 1; i < n-1; i++ {
		if arr[i] != arr[i-1] || arr[i] != arr[i+1] {
			nums = append(nums, arr[i])
		}
	}
	nums = append(nums, arr[n-1])
	for i, v := range nums {
		idxMap[v] = append(idxMap[v], i)
	}
	tar := len(nums) - 1
	steps := make([]int, n, n)
	visited := map[int]struct{}{0: {}}
	q := newQueue()
	q.append(0)
	for len(q) > 0 {
		curID := q.popleft()
		if curID == tar {
			return steps[curID]
		}
		v := nums[curID]
		for _, i := range idxMap[v] {
			if _, ok := visited[i]; !ok {
				visited[i] = struct{}{}
				q.append(i)
				steps[i] = steps[curID] + 1
			}
		}
		if _, ok := visited[curID+1]; !ok && curID+1 <= tar {
			q.append(curID + 1)
			visited[curID+1] = struct{}{}
			steps[curID+1] = steps[curID] + 1
		}
		if _, ok := visited[curID-1]; !ok && curID-1 >= 0 {
			q.append(curID - 1)
			visited[curID-1] = struct{}{}
			steps[curID-1] = steps[curID] + 1
		}
	}
	return -1
}

func main() {
	arr := []int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}
	r := minJumps(arr)
	fmt.Println(r)
}
