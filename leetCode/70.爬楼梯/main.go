package main

import "fmt"

// 当前 i = [i -1] + [i - 2]
func climbStairs(n int) int {
	pre := 1
	cur := 2
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		for k := 3; k <= n; k++ {
			pre, cur = cur, pre+cur
		}
	}
	return cur
}

func main() {
	p := 20
	r := climbStairs(p)
	fmt.Printf("输入:%v   输出:%v\n", p, r)

}
