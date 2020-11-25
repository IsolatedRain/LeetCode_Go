package main

import "fmt"

// BFS 逐层外扩, 检查有无新增.
func orangesRotting(grid [][]int) int {
	dir4 := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	row, col := len(grid), len(grid[0])

	// 存放开始时, 腐烂橘子的坐标
	q := [][]int{}
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if grid[r][c] == 2 {
				q = append(q, []int{r, c})
			}
		}
	}

	days := 0
	for len(q) > 0 {
		nq := [][]int{}
		for _, c := range q {
			for _, m := range dir4 {
				nx := c[0] + m[0]
				ny := c[1] + m[1]
				if nx >= 0 && nx < row && ny >= 0 && ny < col {
					if grid[nx][ny] == 1 {
						grid[nx][ny] = 2
						nq = append(nq, []int{nx, ny})
					}
				}
			}
		}
		q = nq
		// 如果有新的感染,
		// 天数 ++
		if len(q) > 0 {
			days++
		}
	}

	// 有没感染到的桔子, 返回-1
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if grid[r][c] == 1 {
				return -1
			}
		}
	}

	return days
}

func main() {
	p := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	fmt.Printf("输入: %v\n", p)
	r := orangesRotting(p)
	fmt.Printf("输出: %v\n", r)

}
