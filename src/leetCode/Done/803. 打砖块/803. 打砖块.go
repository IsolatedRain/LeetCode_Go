package main

import "fmt"

func hitBricks(grid [][]int, hits [][]int) []int {
	row := len(grid)
	col := len(grid[0])
	root := row*col + 1         // 当作 顶部连通块的根
	p := make([]int, root+2)    // 连通块的父节点
	size := make([]int, root+2) // 每个连通块的大小
	for i := range p {          // 初始化 根节点和连通块大小为1
		p[i] = i
		size[i] = 1
	}
	size[root] = 0

	var getRoot func(x int) int
	getRoot = func(x int) int {
		if p[x] != p[p[x]] {
			p[x] = getRoot(p[x])
		}
		return p[x]
	}

	var union func(x, y int) // 合并连通块，同时合并大小
	union = func(x, y int) {
		xRoot := getRoot(x)
		yRoot := getRoot(y)
		if xRoot == yRoot {
			return
		}
		if xRoot < yRoot {
			p[xRoot] = yRoot
			size[yRoot] += size[xRoot]
			size[xRoot] = 0
		} else {
			p[yRoot] = xRoot
			size[xRoot] += size[yRoot]
			size[yRoot] = 0
		}
	}

	status := make([][]int, row)
	for i := range status {
		status[i] = append(status[i], grid[i]...)
	}
	for _, pos := range hits {
		status[pos[0]][pos[1]] = 0
	}

	for r := 0; r < row; r++ { // 计算最终状态的连通块及其大小
		for c := 0; c < col; c++ {
			if status[r][c] != 1 {
				continue
			}
			if r == 0 {
				union(c, root)
			} else {
				if r > 0 && status[r-1][c] == 1 {
					union((r-1)*col+c, r*col+c)
				}
				if c > 0 && status[r][c-1] == 1 {
					union(r*col+c, r*col+c-1)
				}
			}
		}
	}

	res := make([]int, len(hits))
	dir4 := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for i := len(hits) - 1; i >= 0; i-- { // 倒序填充砖块，合并连通块，检查根节点增加了多少.
		r, c := hits[i][0], hits[i][1]
		preRootSize := size[root]
		if grid[r][c] == 0 {
			continue
		}
		if r == 0 {
			union(c, root)
		}
		for _, v := range dir4 {
			newR := r + v[0]
			newC := c + v[1]
			if 0 <= newR && newR < row &&
				0 <= newC && newC < col &&
				status[newR][newC] == 1 {
				union(r*col+c, newR*col+newC)
			}
		}
		if curAdd := size[root] - preRootSize - 1; curAdd > 0 {
			res[i] = curAdd
		}
		status[r][c] = 1
	}

	return res
}

func main() {
	grid := [][]int{{1}, {1}, {1}, {1}, {1}}
	hits := [][]int{{3, 0}, {4, 0}, {1, 0}, {2, 0}, {0, 0}}
	// grid := [][]int{{1, 1, 0, 1, 0}, {1, 1, 0, 1, 1}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}
	// hits := [][]int{{5, 1}, {1, 3}}
	// grid := [][]int{{1, 0, 0, 0}, {1, 1, 1, 0}}
	// hits := [][]int{{1, 0}}
	// grid := [][]int{{0, 1, 1, 1, 1}, {1, 1, 1, 1, 0}, {1, 1, 1, 1, 0}, {0, 0, 1, 1, 0}, {0, 0, 1, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}
	// hits := [][]int{{6, 0}, {1, 0}, {4, 3}, {1, 2}}
	fmt.Println(hitBricks(grid, hits))
}
