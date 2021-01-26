package main

import "fmt"

func numEquivDominoPairs(dominoes [][]int) int {
	d := map[[2]int]int{}
	for _, v := range dominoes {
		x, y := v[0], v[1]
		if x > y {
			x, y = y, x
		}
		d[[2]int{x, y}]++
	}
	fmt.Println(d)
	var calComb func(n, m int) int
	calComb = func(n, m int) int {
		start := n - m + 1
		cur := 1
		for start <= n {
			cur *= start
			start++
			for cur%m == 0 && m != 1 {
				cur /= m
				m--
			}
		}
		t := 1
		if m != 1 {
			for m > 1 {
				t *= m
				m--
			}
		}
		return cur / t
	}
	res := 0
	for _, v := range d {
		res += calComb(v, 2)
	}
	return res
}

func main() {
	dominoes := [][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}
	fmt.Println(numEquivDominoPairs(dominoes))

}
