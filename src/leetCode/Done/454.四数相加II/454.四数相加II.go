package main

import "fmt"

func fourSumCount(A []int, B []int, C []int, D []int) int {
	mapAB := map[int]int{}
	for _, a := range A {
		for _, b := range B {
			mapAB[a+b]++
		}
	}
	cnt := 0
	for _, c := range C {
		for _, d := range D {
			k := (c + d) * (-1)
			if v, ok := mapAB[k]; ok {
				cnt += v
			}
		}
	}
	return cnt
}

func main() {
	A := []int{1, 2}
	B := []int{-2, -1}
	C := []int{-1, 2}
	D := []int{0, 2}
	r := fourSumCount(A, B, C, D)
	fmt.Printf("%v\n", r)

}
