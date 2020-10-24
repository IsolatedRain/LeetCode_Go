package leetCode

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
