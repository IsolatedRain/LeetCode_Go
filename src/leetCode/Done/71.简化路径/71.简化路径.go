package main

import (
	"fmt"
	// "path/filepath"
	"strings"
)

func simplifyPath(path string) string {
	// return filepath.Clean(path)
	p := strings.Split(path, "/")
	res := []string{}
	for i := range p {
		if isDir(p[i]) {
			res = append(res, p[i])
		} else {
			if p[i] == ".." {
				if len(res) >= 1 {
					res = res[:len(res)-1]
				} else {
					res = res[:0]
				}
			}
		}
	}
	return "/" + strings.Join(res, "/")
}

func isDir(s string) bool {
	if s == "." || s == ".." || s == "" {
		return false
	}
	return true
}

func main() {
	// path := "/a/../../b/../c//.//"
	// path := "/a//b////c/d//././/.."
	// path := "/home//foo/"
	path := "/..."
	r := simplifyPath(path)
	fmt.Println(r)
}
