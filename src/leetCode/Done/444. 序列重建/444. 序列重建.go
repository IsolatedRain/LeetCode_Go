package main

import "fmt"

type set struct {
	dict map[int]struct{}
}

func (s *set) add(v int) {
	s.dict[v] = struct{}{}
}

func (s *set) discard(v int) {
	if _, ok := s.dict[v]; ok {
		delete(s.dict, v)
	}
}
func newSet() set {
	return set{map[int]struct{}{}}
}

func sequenceReconstruction(org []int, seqs [][]int) bool {
	n := len(org)
	graph := make([][2]set, n+1)
	s := newSet()
	for _, seq := range seqs {
		for _, v := range seq {
			s.add(v)
		}
	}
	if len(s.dict) != n {
		return false
	}
	for k := 1; k < n+1; k++ {
		if _, ok := s.dict[k]; !ok {
			return false
		}
	}
	for i := range graph {
		graph[i][0] = newSet()
		graph[i][1] = newSet()
	}
	for _, seq := range seqs {
		for i := 1; i < len(seq); i++ {
			graph[seq[i-1]][1].add(seq[i])
			graph[seq[i]][0].add(seq[i-1])
		}
	}
	q := newSet()
	for i := 1; i < n+1; i++ {
		d := graph[i]
		if cross(d[0], d[1]) {
			return false
		}
		if len(d[0].dict) == 0 {
			q.add(i)
		}
	}
	if len(q.dict) == 0 || len(q.dict) > 1 {
		return false
	}
	arr := []int{}
	for len(q.dict) > 0 {
		nxtQ := newSet()
		for k := range q.dict {
			arr = append(arr, k)
			for p := range graph[k][1].dict {
				if p == 0 {
					continue
				}
				graph[p][0].discard(k)
				if len(graph[p][0].dict) == 0 {
					nxtQ.add(p)
				}
			}
		}
		q = nxtQ
		if len(q.dict) > 1 {
			return false
		}
	}
	return equal(arr, org)
}
func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
func cross(x, y set) bool {
	for k := range x.dict {
		if _, ok := y.dict[k]; ok {
			return true
		}
	}
	return false
}

func main() {
	// org := []int{4, 1, 5, 2, 6, 3}
	// seqs := [][]int{{5, 2, 6, 3}, {4, 1, 5, 2}}
	org := []int{1}
	seqs := [][]int{{2}}

	fmt.Println(sequenceReconstruction(org, seqs))
}
