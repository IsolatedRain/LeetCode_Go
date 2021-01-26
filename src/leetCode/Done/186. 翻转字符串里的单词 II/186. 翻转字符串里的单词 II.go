package main

func reverseWords(s []byte) {
	reverse(s)
	s = append(s, ' ')
	L := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			reverse(s[L:i])
			L = i + 1
		}
	}
	return
}
func reverse(s []byte) []byte {
	L, R := 0, len(s)-1
	for L < R {
		s[L], s[R] = s[R], s[L]
		L++
		R--
	}
	return s
}

func main() {
	s := []byte{'t', 'h', 'e', ' ', 's', 'k', 'y', ' ', 'i', 's', ' ', 'b', 'l', 'u', 'e'}
	reverseWords(s)
}
