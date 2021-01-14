package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)
	L, R := 0, len(words)-1
	for L < R {
		words[L], words[R] = words[R], words[L]
		L++
		R--
	}
	return strings.Join(words, " ")
}

// L, R := 0, len(s)-1
// for L < R && s[L] == ' ' {
// 	L++
// }
// for R > L && s[R] == ' ' {
// 	R--
// }
// words := []string{}
// word := []byte{}
// for L <= R {
// 	if s[L] != ' ' {
// 		word = append(word, s[L])
// 	} else if len(word) > 0 {
// 		words = append(words, string(word))
// 		word = []byte{}
// 	}
// 	L++
// }
// words = append(words, string(word))
// L, R = 0, len(words)-1
// for L < R {
// 	words[L], words[R] = words[R], words[L]
// 	L++
// 	R--
// }
// return strings.Join(words, " ")
// }

func main() {
	s := "    the sky is blue  "
	fmt.Println(reverseWords(s))
}
