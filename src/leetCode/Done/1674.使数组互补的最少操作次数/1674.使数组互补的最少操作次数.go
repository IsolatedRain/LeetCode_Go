package main

import "fmt"

// 数对 去更新 差分数组
// 差分数组 表示操作的次数
// 2-> L+ 1 , 操作2次
// L + 1 -> R+limit 操作1次
// L + R -> 操作0次
// R + limit + 1 -> 操作2次
func minMoves(nums []int, limit int) int {
	n := len(nums)
	diff := make([]int, 2*limit+2, 2*limit+2)
	for i := 0; i < n>>1; i++ {
		L := min(nums[i], nums[n-i-1])
		R := max(nums[i], nums[n-i-1])
		fmt.Println(L, R)
		diff[2] += 2
		diff[L+1] -= 2
		diff[L+1]++
		diff[R+limit+1]--
		diff[L+R]--
		diff[L+R+1]++
		diff[R+limit+1] += 2
	}
	res := n
	cur := 0
	for i := range diff[2 : len(diff)-1] {
		cur += diff[i+2]
		if cur < res {
			res = cur
		}
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	// nums := []int{1, 2, 2, 1}
	nums := []int{1, 2, 3, 4}
	limit := 4
	r := minMoves(nums, limit)
	fmt.Println(r)

}
