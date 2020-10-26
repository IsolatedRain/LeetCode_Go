package leetcode

// 思路: 保存 滑动窗口 L -> curIndex的大小,
//出现重复字母的时候, 更新 L 位置到 上次出现重复字母的右边.
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	d := make(map[rune]int)
	res := 0
	L := 0
	for i, k := range s {
		lastIdx, exist := d[k]
		if exist && lastIdx >= L {
			L = lastIdx + 1
		}
		res = max(res, i-L)
		d[k] = i
	}
	return res + 1
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
