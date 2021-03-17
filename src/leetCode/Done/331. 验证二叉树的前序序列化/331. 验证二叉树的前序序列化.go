package main

import (
	"fmt"
	"regexp"
)

// 不断摘除叶子节点， 即 替换 叶子节点为 #
func isValidSerialization(preorder string) bool {
	s := []byte(preorder)
	for {
		n := len(s)
		re, err := regexp.Compile("[0-9]+,#,#")
		if err != nil {
			break
		}
		s = re.ReplaceAll(s, []byte{'#'})
		if len(s) == n {
			break
		}
	}
	return len(s) == 1 && s[0] == '#'
}

func main() {
	fmt.Println(isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"))
}
