package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	n := len(arr)
	cnt := make(map[int]int, n)
	for i := 0; i < n; i++ {
		cnt[arr[i]]++
	}
	d := make(map[int]int, len(cnt))
	for _, v := range cnt {
		_, exsit := d[v]
		if exsit {
			return false
		}
		d[v]++
	}
	return true
}

func main() {
	para := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	r := uniqueOccurrences(para)
	fmt.Printf("%v\n", r)

}
