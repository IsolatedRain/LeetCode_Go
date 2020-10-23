package leetCode

import "fmt"

func twoSum(nums []int, target int) []int {
	d := make(map[int]int, 0)
	for i, num := range nums {
		j, exist := d[target-num]
		if exist {
			return []int{j, i}
		}
		d[num] = i
	}
	return []int{}
}

func main() {
	// test case
	// one
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
