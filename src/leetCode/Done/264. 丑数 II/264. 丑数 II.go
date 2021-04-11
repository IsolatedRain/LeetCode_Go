package main

import "fmt"

func nthUglyNumber(n int) int {
	res := []int{1}
	p2, p3, p5 := 0, 0, 0
	for n > 1 {
		curMin := min(2*res[p2], 3*res[p3], 5*res[p5])
		if 2*res[p2] == curMin {
			p2++
		}
		if 3*res[p3] == curMin {
			p3++
		}
		if 5*res[p5] == curMin {
			p5++
		}
		res = append(res, curMin)
		n--
	}
	return res[len(res)-1]
}
func min(arr ...int) int {
	r := arr[0]
	for _, v := range arr[1:] {
		if v < r {
			r = v
		}
	}
	return r
}

func main() {
	fmt.Println(nthUglyNumber(11))
}
