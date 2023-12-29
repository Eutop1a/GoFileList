package main

import (
	"fmt"
)

func rotateRange(s []int, k int) []int {
	n := len(s)
	res := make([]int, n)
	for i, val := range s {
		res[(i+k)%n] = val
	}
	return res
}

func rotate(s []int, k int) {
	k = k % len(s)
	reverse(s)
	reverse(s[k:])
	reverse(s[:k])
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	rotate(s, 3)
	fmt.Println(s)
	fmt.Println(rotateRange(s, 3))
}
