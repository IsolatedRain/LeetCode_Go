package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	k := 1
	res := "1"
	for k < n {
		k++
		pre := res[0]
		var tmp strings.Builder
		cnt := 1
		for i := 1; i < len(res); i++ {
			if res[i] == pre {
				cnt++
			} else {
				tmp.WriteString(strconv.Itoa(cnt))
				tmp.WriteByte(pre)
				pre = res[i]
				cnt = 1
			}
		}
		tmp.WriteString(strconv.Itoa(cnt))
		tmp.WriteByte(pre)
		res = tmp.String()
	}
	return res

}

type test struct {
	para int
	res  string
}

func main() {
	testCases := []test{
		{
			1,
			"1",
		},
		{
			5,
			"111221",
		},
		{
			10,
			"13211311123113112211",
		},
	}
	for _, p := range testCases {
		curPara := p.para
		correctRes := p.res
		res := countAndSay(curPara)
		fmt.Printf("输入:%v\n输出:  %v\n预期:  %v\n", curPara, res, correctRes)
	}
}
