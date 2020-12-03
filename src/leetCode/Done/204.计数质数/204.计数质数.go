package main

import "fmt"

func countPrimes(n int) int {
	mark := make([]bool, n)
	primes := []int{}
	for i := 2; i < n; i++ {
		if !mark[i] {
			primes = append(primes, i)
		}
		for _, p := range primes {
			if p*i >= n {
				break
			}
			mark[p*i] = true
			if i%p == 0 {
				break
			}
		}

	}
	return len(primes)
}

// func countPrimes(n int) int {
// 	arr := make([]bool, n)
// 	cnt := 0
// 	for i := 2; i < n; i++ {
// 		if arr[i] {
// 			continue
// 		}
// 		cnt++
// 		for j := i + i; j < n; j += i {
// 			arr[j] = true
// 		}
// 	}
// 	return cnt
// }

func main() {
	n := 100
	r := countPrimes(n)
	fmt.Println(r)
}
