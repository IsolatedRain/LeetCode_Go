package main

import (
	"fmt"
	"math/big"
)

func waysToFillArray(queries [][]int) []*big.Int {
	mark := make([]bool, 10002)
	primes := []int{}
	mod := int64(1000000000 + 7)
	for i := 2; i < 10002; i++ {
		if !mark[i] {
			primes = append(primes, i)
		}
		for _, p := range primes {
			if p*i >= 10002 {
				break
			}
			mark[p*i] = true
			if i%p == 0 {
				break
			}
		}
	}

	var calEX func(num int) map[int]int
	calEX = func(num int) map[int]int {
		r := map[int]int{}
		for _, p := range primes {
			for num%p == 0 && num > 1 {
				r[p]++
				num /= p
			}
		}
		return r
	}

	res := []*big.Int{}
	for _, q := range queries {
		n, k := q[0], q[1]
		cur := big.NewInt(1)
		c := calEX(k)
		fmt.Println(c)
		for _, v := range c {
			cur = cur.Mul(cur, calComb(n, v))
		}
		cur.Mod(cur, big.NewInt(mod))
		res = append(res, cur)
	}
	return res
}
func calComb(n, r int) *big.Int {
	end := int64(n + r - 1)
	cur := big.NewInt(1)
	var i int64
	for i = int64(n); i <= end; i++ {
		cur.Mul(cur, big.NewInt(i))
	}
	curR := 1
	for i := 2; i < r; i++ {
		curR *= i
	}
	return cur.Div(cur, big.NewInt(int64(curR)))
}
func main() {
	queries := [][]int{{3360, 1536}}
	fmt.Println(waysToFillArray(queries))
}
