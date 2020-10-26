package leetcode

func smallerNumbersThanCurrent(nums []int) []int {
	max_num := max(nums)
	count := make([]int, max_num+1)
	for _, v := range nums {
		count[v]++
	}
	pre_sum := make([]int, max_num+1)
	pre_sum[0] = count[0]
	for i := 0; i < max_num; i++ {
		pre_sum[i+1] = count[i+1] + pre_sum[i]
	}
	res := make([]int, len(nums))
	for i, v := range nums {
		if v > 0 {
			res[i] = pre_sum[v-1]
		}
	}

	return res
}
func max(arr []int) int {
	n := len(arr)
	if n == 1 {
		return arr[0]
	}
	x := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] > x {
			x = arr[i]
		}
	}
	return x
}
