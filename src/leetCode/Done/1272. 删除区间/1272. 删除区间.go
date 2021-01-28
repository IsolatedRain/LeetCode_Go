package main

import "fmt"

func removeInterval(intervals [][]int, toBeRemoved []int) [][]int {
	res := [][]int{}
	L, R := toBeRemoved[0], toBeRemoved[1]
	for _, i := range intervals {
		if i[0] > R || i[1] < L {
			res = append(res, i)
		} else {
			if i[0] < L && L < i[1] {
				res = append(res, []int{i[0], L})
			}
			if i[0] < R && R < i[1] {
				res = append(res, []int{R, i[1]})
			}
		}
	}
	return res
}
func main() {
	fmt.Println(removeInterval([][]int{{0, 2}, {3, 4}, {5, 7}}, []int{1, 6}))
}
