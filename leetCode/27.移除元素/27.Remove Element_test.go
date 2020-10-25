package leetCode

import (
	"fmt"
	"testing"
)

type test struct {
	nums []int
	val  int
	res  int
}

func Test(t *testing.T) {
	test_cases := []test{
		{
			[]int{3, 2, 2, 3},
			3,
			2,
		},
		{
			[]int{0, 1, 2, 2, 3, 0, 4, 2},
			2,
			5,
		},
		{
			[]int{2, 2, 2},
			2,
			0,
		},
		{
			[]int{},
			0,
			0,
		},
	}
	fmt.Println("------------------------移除元素---------------------------")
	for _, p := range test_cases {
		cur_para := p.nums
		cur_val := p.val
		correct_res := p.res
		cur_res := removeElement(cur_para, cur_val)
		fmt.Printf("输入:%v,    输出:%v,     预期结果:%v    \n", cur_para, cur_res, correct_res)
		fmt.Printf("结果正确:  %v\n", correct_res == cur_res)
	}
}
