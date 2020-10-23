package leetCode

import "fmt"

func romanToInt(s string) int {
	dic := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000}
	res := 0
	pre := 0
	for i := len(s) - 1; i >= 0; i-- {
		cur := dic[s[i]]
		if cur >= pre {
			res += cur
		} else if cur < pre {
			res -= cur
		}
		pre = cur
	}
	return res
}

func main() {
	// 测试用例
	s := "LVIII"
	s1 := "MCMXCIV"
	fmt.Println(romanToInt(s))
	fmt.Println(romanToInt(s1))
}
