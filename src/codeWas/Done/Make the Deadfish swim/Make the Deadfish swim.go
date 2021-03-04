//https://www.codewars.com/kata/51e0007c1f9378fa810002a9/train/go
package main

import "fmt"

// Parse 计算和输出
func Parse(data string) []int {
	res := []int{}
	curValue := 0
	for i := range data {
		switch data[i] {
		case 'i':
			curValue++
		case 'd':
			curValue--
		case 's':
			curValue *= curValue
		case 'o':
			res = append(res, curValue)
		}
	}
	return res
}

func main() {
	fmt.Println(Parse("isoisoiso"), []int{1, 4, 25})
}
