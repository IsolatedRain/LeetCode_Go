package main

import (
	"fmt"
	"strconv"
)

func similarRGB(color string) string {
	return "#" + format(color[1:3]) + format(color[3:5]) + format(color[5:])
}

func format(s string) string {
	// 转为10进制数
	n, _ := strconv.ParseInt(s, 16, 32)
	// 检查余数 是否大于 8
	// 超过8， 则 + 1的数更接近。
	// 17 的16进制表示为 0x11
	r, q := n%17, n/17
	if r > 8 {
		q++
	}
	return fmt.Sprintf("%02x", 17*q)
}
func main() {
	color := "#09f166"
	r := similarRGB(color)
	fmt.Println("输出：", r)
}
