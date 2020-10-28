package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	cnt := map[int]int{}
	for _, k := range arr {
		cnt[k]++
	}

	d := map[int]bool{}
	for _, c := range cnt {
		if d[c] {
			return false
		}
		d[c] = true
	}
	return true
}

func main() {
	para := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	r := uniqueOccurrences(para)
	fmt.Printf("%v\n", r)

}
