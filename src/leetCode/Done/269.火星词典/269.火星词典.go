package main

import (
	"fmt"
)

func alienOrder(words []string) string {
	inDegree := map[byte]int{}      // 存字母的入度, 即 在它前面的字母个数
	suf := map[byte][]byte{}        // 存字母的 后继,必须在它后面的那些字母.
	paired := map[string]struct{}{} // 存已经比较过的 字母对 如 [wz]:{} w在z前面,已经比较过了. 如果出现 zw, 就存在环, ""

	lenWords := len(words)
	for i := range words {
		for j := i + 1; j < lenWords; j++ {
			curLen := min(len(words[i]), len(words[j]))
			idx := 0 // 逐位比较到结束(较短的那个单词的结尾)
			for idx < curLen {
				if words[i][idx] != words[j][idx] {
					// 遇到不同, 就退出比较
					break
				}
				idx++
			}

			if idx == curLen {
				//  "abc", "ab" 到了短的结尾, 实际应为 "ab", "abc"
				// 矛盾了, return ""
				if len(words[i]) > curLen {
					return ""
				}
				// 两个 字符完全一样, 就比较下一对.
				continue
			}
			// 正向 / 反向
			tmpK := string(words[i][idx]) + string(words[j][idx])
			tmpKr := string(words[j][idx]) + string(words[i][idx])
			// 如果反向 也比较过, 已存, 就是说 成环,
			// return ""
			if _, ok := paired[tmpKr]; ok {
				return ""
			}
			// 已经比较过, 就跳过
			if _, ok := paired[tmpK]; ok {
				continue
			}
			paired[tmpK] = struct{}{}
			// 加入后继
			suf[words[i][idx]] = append(suf[words[i][idx]], words[j][idx])
			inDegree[words[j][idx]]++ // 入度数 +1
		}
	}

	start := map[byte]struct{}{}
	for _, word := range words {
		for i := range word {
			// 没有存在 入度字典里, 即 入度为0,
			// 作为 拓扑排序的起点
			if _, ok := inDegree[word[i]]; !ok {
				start[word[i]] = struct{}{}
			}
		}
	}

	res := []byte{}
	for len(start) > 0 {
		nextStart := map[byte]struct{}{}
		for c := range start {
			res = append(res, c)
			for _, s := range suf[c] {
				inDegree[s]--
				// 入度为0, 可以作为 下一次排序的起点(逐层BFS)
				if inDegree[s] == 0 {
					nextStart[s] = struct{}{}
				}
			}
		}
		start = nextStart
	}

	return string(res)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	words := []string{
		"wrt",
		"wrf",
		"er",
		"ett",
		"rftt"}
	// words := []string{"abc", "ab"}
	fmt.Printf("输入: %v\n", words)
	r := alienOrder(words)
	fmt.Printf("输出: %v\n", r)

}
