package main

import (
	"fmt"
	"strings"
)

func smallestEquivalentString(A string, B string, S string) string {
	p := make(map[byte]byte)
	var getRoot func(x byte) byte
	getRoot = func(x byte) byte {
		if _, ok := p[x]; !ok {
			p[x] = x
		}
		if p[x] != p[p[x]] {
			p[x] = getRoot(p[x])
		}
		return p[x]
	}
	var union func(x, y byte)
	union = func(x, y byte) {
		xRoot := getRoot(x)
		yRoot := getRoot(y)
		if xRoot > yRoot {
			p[xRoot] = yRoot
		} else {
			p[yRoot] = xRoot
		}
	}

	n := len(A)
	for i := 0; i < n; i++ {
		xR := getRoot(A[i])
		yR := getRoot(B[i])
		if xR != yR {
			union(A[i], B[i])
		}
	}
	fmt.Printf("%v\n", p)
	var res strings.Builder
	for i := 0; i < len(S); i++ {
		sRoot := getRoot(S[i])
		res.WriteByte(sRoot)
	}
	return res.String()
}

func main() {
	A := "parker"
	B := "morris"
	S := "parser"
	fmt.Printf("输入: %v %v %v\n", A, B, S)
	r := smallestEquivalentString(A, B, S)
	fmt.Printf("输出: %v\n", r)
}
