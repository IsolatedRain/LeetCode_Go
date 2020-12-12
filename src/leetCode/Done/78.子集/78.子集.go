package main

import "fmt"

func subsets(nums []int) (res [][]int) {
	end := 1 << len(nums)
	for mask := 0; mask < end; mask++ {
		cur := []int{}
		for i, v := range nums {
			if (mask & (1 << i)) != 0 {
				cur = append(cur, v)
			}
		}
		res = append(res, cur)
	}
	return
}

func main() {
	nums := []int{1, 2, 3}
	r := subsets(nums)
	fmt.Println(r)

}
