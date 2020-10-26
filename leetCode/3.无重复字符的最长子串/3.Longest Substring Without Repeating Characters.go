package leetcode

func lengthOfLongestSubstring(s string) int {
	d := make(map[rune]int)
	res := 0
	for i, k := range s {
		lastIdx, exist := d[k]
		if exist {
			res = max(res, i-lastIdx)
			d[k] = i
		} else {
			d[k] = i
		}
	}
	return res
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
