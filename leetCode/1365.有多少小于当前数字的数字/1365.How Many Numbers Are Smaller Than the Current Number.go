package leetcode

func smallerNumbersThanCurrent(nums []int) []int {
	maxNum := max(nums)
	count := make([]int, maxNum+1)
	for _, v := range nums {
		count[v]++
	}
	preSum := make([]int, maxNum+1)
	preSum[0] = count[0]
	for i := 0; i < maxNum; i++ {
		preSum[i+1] = count[i+1] + preSum[i]
	}
	res := make([]int, len(nums))
	for i, v := range nums {
		if v > 0 {
			res[i] = preSum[v-1]
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
