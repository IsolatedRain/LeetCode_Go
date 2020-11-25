package main

import "fmt"

func toLowerCase(str string) string {
	s := []byte(str)
	fmt.Printf("%v\n", s)
	for i, v := range s {
		if v >= 'A' && v <= 'Z' {
			s[i] += 32
		}
	}
	return string(s)
	// return strings.ToLower(str)
}

func main() {
	p := "LOVELYgirL"
	r := toLowerCase(p)
	fmt.Printf("输入: %v\n 输出: %v\n", p, r)
}
