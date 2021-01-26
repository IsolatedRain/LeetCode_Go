package main

import "fmt"

func addToArrayForm(A []int, K int) []int {
	reverse(A)
	B := []int{}
	for K > 0 {
		B = append(B, K%10)
		K /= 10
	}
	carry := 0
	res := []int{}
	if len(A) < len(B) {
		A, B = B, A
	}
	i := 0
	for i < len(B) {
		curDigit := A[i] + B[i] + carry
		res = append(res, curDigit%10)
		carry = curDigit / 10
		i++
	}
	for i < len(A) {
		curDigit := A[i] + carry
		res = append(res, curDigit%10)
		carry = curDigit / 10
		i++
	}
	if carry != 0 {
		res = append(res, 1)
	}
	return reverse(res)
}
func reverse(arr []int) []int {
	L, R := 0, len(arr)-1
	for L < R {
		arr[L], arr[R] = arr[R], arr[L]
		L++
		R--
	}
	return arr
}

func main() {
	A := []int{1, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	K := 12
	fmt.Println(addToArrayForm(A, K))
}
