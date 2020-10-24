package leetCode

import "math"

func longestCommonPrefix(strs []string) string {
	min_s := ""
	cur_len := math.MaxInt32
	for _, s := range strs {
		if len(s) < cur_len {
			min_s = s
			cur_len = len(s)
		}
	}
	if min_s == "" {
		return min_s
	}
	for _, s := range strs {
		for i := 0; i < len(min_s); i++ {
			if min_s[i] != s[i] {
				min_s = min_s[:i]
				break
			}
		}
	}
	return min_s
}
