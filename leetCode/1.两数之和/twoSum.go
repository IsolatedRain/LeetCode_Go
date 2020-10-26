package leetcode

// 思路: i + j = target
//  	 target - i = index 存入 字典,
//		如果 遇到 j 在字典里, 表示  前面已经出现 j = target - i,
//		其索引 index = map[j] 取出, []int{j, i} 即为结果.
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
