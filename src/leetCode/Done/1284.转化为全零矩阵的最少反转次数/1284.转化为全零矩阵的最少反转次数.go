package main

import "fmt"

func minFlips(mat [][]int) int {
	row := len(mat)
	col := len(mat[0])
	mask := 0
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			mask = mask << 1
			mask |= mat[r][c]
		}
	}

	// mask转为矩阵
	var mask2mat func(mask int) [][]int
	mask2mat = func(mask int) [][]int {
		curMat := make([][]int, row)
		for r := 0; r < row; r++ {
			curMat[r] = make([]int, col)
		}
		for r := row - 1; r > -1; r-- {
			for c := col - 1; c > -1; c-- {
				curMat[r][c] = mask & 1
				mask >>= 1
			}
		}
		return curMat
	}
	// 翻转后 返回 mask
	var flip func(x, y int, curMat [][]int) int
	flip = func(x, y int, curMat [][]int) int {
		curMask := 0
		for r := 0; r < row; r++ {
			for c := 0; c < col; c++ {
				curMask <<= 1
				xorCurMat := curMat[r][c] ^ 1
				if (r == x && (c == y || c == y-1 || c == y+1)) || (c == y && (r == x || r == x-1 || r == x+1)) {
					curMask |= xorCurMat
				} else {
					curMask |= curMat[r][c]
				}
			}
		}
		return curMask
	}

	checked := map[int]bool{}
	q := []int{mask}
	steps := 0
	for len(q) > 0 {
		n := len(q)
		for _, m := range q {
			if m == 0 {
				return steps
			}
			curMat := mask2mat(m)
			for r := 0; r < row; r++ {
				for c := 0; c < col; c++ {
					nxtMask := flip(r, c, curMat)
					if !checked[nxtMask] {
						checked[nxtMask] = true
						q = append(q, nxtMask)
					}
				}
			}

		}
		q = q[n:]
		steps++
	}
	return -1
}

func main() {
	mat := [][]int{{1, 1, 1}, {1, 0, 1}, {0, 0, 0}}
	r := minFlips(mat)
	fmt.Println(r)
}
