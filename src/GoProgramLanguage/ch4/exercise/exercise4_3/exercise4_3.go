package main

import "fmt"

func reverse(s *[length]int) {
	n := length - 1
	for i, j := 0, n; i < j; i, j = i+1, j-1 {
		//(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
		s[i], s[j] = s[j], s[i]
	}
}

const length = 6

func main() {
	s := [6]int{1, 2, 3, 4, 5, 6}
	reverse(&s)
	fmt.Println(s)
}
