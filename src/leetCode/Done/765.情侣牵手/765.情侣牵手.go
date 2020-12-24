package main

import "fmt"

// UF 并查集
type UF struct {
	p map[int]int
}

func (u *UF) find(x int) int {
	if u.p[x] != x {
		u.p[x] = u.find(u.p[x])
	}
	return u.p[x]
}

func (u *UF) union(x, y int) {
	xRoot := u.find(x)
	yRoot := u.find(y)
	if xRoot > yRoot {
		u.p[xRoot] = yRoot
	} else {
		u.p[yRoot] = xRoot
	}
}

func minSwapsCouples(row []int) int {
	n := len(row)
	p := map[int]int{}
	for i := 0; i < n; i += 2 {
		p[i] = i
		p[i+1] = i
	}
	u := new(UF)
	u.p = p
	for i := 0; i < n; i += 2 {
		u.union(row[i], row[i+1])
	}
	res := 0
	for i := 0; i < n; i += 2 {
		xRoot := u.find(i)
		if xRoot != i {
			res++
		}
	}
	return res
}

func main() {
	// row := []int{0, 2, 1, 3}
	row := []int{5, 3, 4, 2, 1, 0}
	// row := []int{0, 2, 4, 6, 7, 1, 3, 5}
	r := minSwapsCouples(row)
	fmt.Println(r)
}
