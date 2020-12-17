package main

import (
	"fmt"
)

func maxVacationDays(flights [][]int, days [][]int) int {
	weeks := len(days[0])
	cities := len(flights)
	if weeks == 0 || cities == 0 {
		return 0
	}
	for i := 0; i < cities; i++ {
		flights[i][i] = 1
	}
	last := make([]int, cities, cities)
	for i := 0; i < cities; i++ {
		if flights[0][i] == 1 {
			last[i] = days[i][0]
		} else {
			last[i] = -1
		}
	}
	for week := 1; week < weeks; week++ {
		cur := make([]int, cities, cities)
		for c := 0; c < cities; c++ {
			cur[c] = last[c]
			for from := 0; from < cities; from++ {
				if last[from] == -1 || flights[from][c] == 0 {
					continue
				}
				cur[c] = max(cur[c], days[c][week]+last[from])
			}
		}
		last = cur
	}
	return max(last...)
}

func max(x ...int) int {
	r := x[0]
	for _, v := range x {
		if v > r {
			r = v
		}
	}
	return r
}

func main() {
	flights := [][]int{{0, 1, 1}, {1, 0, 1}, {1, 1, 0}}
	days := [][]int{{1, 3, 1}, {6, 0, 3}, {3, 3, 3}}
	// flights := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	// days := [][]int{{1, 1, 1}, {7, 7, 7}, {7, 7, 7}}
	r := maxVacationDays(flights, days)
	fmt.Println(r)
}
