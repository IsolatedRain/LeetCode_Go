package main

import "fmt"

// x 轴 和 y轴 分开统计 最后是否为0
func judgeCircle(moves string) bool {
	x, y := 0, 0
	for _, char := range moves {
		switch char {
		case 'U':
			y++
		case 'D':
			y--
		case 'L':
			x--
		case 'R':
			x++
		}
	}
	return x == 0 && y == 0
}

func main() {
	p := "UD"
	r := judgeCircle(p)
	fmt.Printf("输入: %v\n输出: %v\n", p, r)
}
