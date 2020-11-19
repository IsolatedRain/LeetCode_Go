package leetcode

func is_same_silice(arr1 []int, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i, v := range arr1 {
		if arr2[i] != v {
			return false
		}
	}
	return true
	// 测试
}
