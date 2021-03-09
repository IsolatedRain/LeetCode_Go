// https://www.codewars.com/kata/5d6eef37f257f8001c886d97/train/go
package main

import "fmt"

// Solver 连续 n 位的和  不超过 q
func Solver(arr []int, n, q int) []int {
	arrLen := len(arr)
	mask := 0  // mark picked num in arr
	path := []int{}

	var dfs func(i, curLen, curSum int) bool
	dfs = func(i, curLen, curSum int) bool {
		// if > limit
		if curSum > q {
			return false
		}
		// if reach end
		if i == arrLen {
			return true
		}
		// if reach limit-Len
		if curLen == n {
			curSum -= path[i-n]  // subtract left-element
			curLen--
		}

		for idx := 0; idx < arrLen; idx++ {
			if mask&(1<<uint(idx)) == 0 {
				mask |= 1 << uint(idx)  // go 1.12 shift use unsigned int
				path = append(path, arr[idx])
				if dfs(i+1, curLen+1, curSum+arr[idx]) {
					return true
				}
				// traceback
				mask ^= 1 << uint(idx)
				path = path[:len(path)-1]
			}
		}
		return false
	}

	dfs(0, 0, 0)
	return path
}

func main() {
	fmt.Println(Solver([]int{3, 5, 7, 1, 6, 8, 2, 4}, 3, 13))
}
