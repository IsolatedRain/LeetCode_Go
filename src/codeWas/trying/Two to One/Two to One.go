// https://www.codewars.com/kata/5656b6906de340bd1b0000ac/train/go
package main

import "fmt"

// TwoToOne 合并二字符串为 不重复字符的字符串
func TwoToOne(s1 string, s2 string) string {
	mark := make([]byte, 26)
	for _, c := range s1 + s2 {
		mark[c-'a'] = 1
	}

	res := ""
	for i := 0; i < 26; i++ {
		if mark[i] != 0 {
			res += string(byte('a' + i))
		}

	}
	return res
}

func main() {
	fmt.Println(TwoToOne("loopingisfunbutdangerous", "lessdangerousthancoding") == "abcdefghilnoprstu")
}
