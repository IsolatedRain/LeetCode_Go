package main

import "fmt"

// 在 2 的整数次幂 i，置1，  然后 +1 +2 ... + i 幂次
// 9 -> 1 0 0 1	 8 + 1
// 8 -> 1 0 0 0
// 7 ->  1 1 1	 4 + 3
// 6 ->  1 1 0	 4 + 2
// 5 ->  1 0 1	 4 + 1
// 4 ->  1 0 0
// 3 ->    1 1	 2 + 1
// 2 ->    1 0
// 1 ->      1
func countBits(num int) []int {
	bit := make([]int, num+1)
	highBit := 0
	for i := 1; i < num+1; i++ {
		if i&(i-1) == 0 {
			bit[i] = 1
			highBit = i
		} else {
			bit[i] = bit[i-highBit] + 1

		}
	}
	return bit
}

func main() {
	fmt.Println(countBits(8))
}
