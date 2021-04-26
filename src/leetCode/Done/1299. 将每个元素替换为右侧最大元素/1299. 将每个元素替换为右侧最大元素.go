// https://leetcode-cn.com/problems/replace-elements-with-greatest-element-on-right-side/
package main

import "fmt"

func replaceElements(arr []int) []int {
	rightMax := -1
	n := len(arr)
	for i := n - 1; i >= 0; i-- {
		tmp := max(rightMax, arr[i])
		arr[i], rightMax = rightMax, tmp
	}
	return arr
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(replaceElements([]int{17, 18, 5, 4, 6, 1}))
}
