package main

import "fmt"

func killProcess(pid []int, ppid []int, kill int) (res []int) {
	graph := map[int][]int{}
	for i := range pid {
		graph[ppid[i]] = append(graph[ppid[i]], pid[i])
	}
	q := []int{kill}
	for len(q) > 0 {
		curLen := len(q)
		for i := 0; i < curLen; i++ {
			res = append(res, q[i])
			if v, ok := graph[q[i]]; ok {
				q = append(q, v...)
			}
		}
		q = q[curLen:]
	}
	return
}

func main() {
	pid := []int{1, 3, 10, 5}
	ppid := []int{3, 0, 5, 3}
	kill := 5
	fmt.Println(killProcess(pid, ppid, kill))
}
