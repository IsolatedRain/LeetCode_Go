package main

func rotate(nums []int, k int) {
	var reverse func(L, R int)
	reverse = func(L, R int) {
		for L < R {
			nums[L], nums[R] = nums[R], nums[L]
			L++
			R--
		}
	}

	k %= len(nums)
	reverse(0, len(nums)-1)
	reverse(0, k-1)
	reverse(k, len(nums)-1)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
}
