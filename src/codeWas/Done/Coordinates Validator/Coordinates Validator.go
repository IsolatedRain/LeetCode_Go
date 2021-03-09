// https://www.codewars.com/kata/5269452810342858ec000951/train/go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// IsValidCoordinates 判断 是否有效的经纬度坐标
func IsValidCoordinates(coordinates string) bool {
	if strings.Contains(coordinates, "e") {
		return false
	}
	s := strings.ReplaceAll(coordinates, ",", "")
	c := strings.Split(s, " ")
	if len(c) != 2 {
		return false
	}
	n1, ok1 := strconv.ParseFloat(c[0], 64)
	n2, ok2 := strconv.ParseFloat(c[1], 64)
	if ok1 != nil || ok2 != nil {
		return false
	}
	fmt.Println(n1, n2)
	if n1 < -90 || n1 > 90 {
		return false
	}
	if n2 < -180 || n2 > 180 {
		return false
	}
	return true
}

func main() {
	// fmt.Println(IsValidCoordinates("24.53525235, 23.45235"))
	fmt.Println(IsValidCoordinates("23.245, 1e1"))
}
