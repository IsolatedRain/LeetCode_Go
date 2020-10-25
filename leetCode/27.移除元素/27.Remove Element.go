package leetCode
// 思路: 左右指针,左指针在0位置, 右指针在数组末尾, 
// 通过交换指针的值, 移动指针.
// 最后检查 左指针位置 是否为 val, 需不需要舍弃.
func removeElement(nums []int, val int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	L := 0
	R := n - 1
	for L < R {
		if nums[L] == val {
			nums[L] = nums[R]
			R -= 1
		} else {
			L += 1
		}
	}
	if nums[L] == val {
		L -= 1
	}
	return L + 1
}
