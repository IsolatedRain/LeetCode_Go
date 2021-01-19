package main

import (
	"fmt"
	"strconv"
	"strings"
)

func toHexspeak(num string) string {
	curNum, _ := strconv.Atoi(num)
	curHex := fmt.Sprintf("%X", curNum)
	for _, b := range curHex {
		if '1' < b && b <= '9' {
			return "ERROR"
		}
	}
	return strings.NewReplacer("0", "O", "1", "I").Replace(curHex)
	// digits := map[int]byte{10: 'A', 11: 'B', 12: 'C', 13: 'D', 14: 'E', 15: 'F', 1: 'I', 0: 'O'}
	// for i := 2; i < 10; i++ {
	// 	digits[i] = byte(i)
	// }
	// res := []byte{}
	// curNum, _ := strconv.Atoi(num)
	// for curNum > 0 {
	// 	curByte := curNum % 16
	// 	if curByte > 1 && curByte < 10 {
	// 		return "ERROR"
	// 	}
	// 	res = append(res, digits[curByte])
	// 	curNum /= 16
	// }
	// L, R := 0, len(res)-1
	// for L < R {
	// 	res[L], res[R] = res[R], res[L]
	// 	L++
	// 	R--
	// }
	// return string(res)
}

func main() {
	num := "257"
	fmt.Println(toHexspeak(num))
}
