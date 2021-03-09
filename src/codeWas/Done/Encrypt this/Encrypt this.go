// https://www.codewars.com/kata/5848565e273af816fb000449/train/go
package main

import (
	"fmt"
	"strings"
)

// EncryptThis 加密
func EncryptThis(text string) string {
	encryptWord := func(word string) string {
		if len(word) == 1 {
			return fmt.Sprintf("%d", word[0])
		}
		w := []byte(word[1:])
		w[0], w[len(w)-1] = w[len(w)-1], w[0]
		return fmt.Sprintf("%d", word[0]) + string(w)
	}
	words := strings.Split(text, " ")
	res := []string{}
	for _, word := range words {
		res = append(res, encryptWord(word))
	}
	return strings.Join(res, " ")
}

func main() {
	// fmt.Println(EncryptThis("The more he saw the less he spoke") == "84eh 109ero 104e 115wa 116eh 108sse 104e 115eokp")
	fmt.Println(EncryptThis("A wise old owl lived in an oak"))
}
