// https://www.codewars.com/kata/51b6249c4612257ac0000005/train/go
package main

import (
	"fmt"
)

// Decode 罗马数字转阿拉伯数字
func Decode(roman string) int {
	d := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000}
	n := len(roman)
	res := d[roman[n-1]]
	for i := n - 2; i >= 0; i-- {
		if d[roman[i]] < d[roman[i+1]] {
			res -= d[roman[i]]
		} else {
			res += d[roman[i]]
		}
	}
	return res
}

func main() {
	fmt.Println(Decode("MDCLXVI"))

}
