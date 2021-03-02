package main

import (
	"fmt"
	"strings"
)

// MakeUpperCase 改为大写
func MakeUpperCase(str string) string {
	s := strings.ToUpper(str)
	return s
}
func main() {
	fmt.Println(MakeUpperCase("1,2,3 hello world!"))
}
