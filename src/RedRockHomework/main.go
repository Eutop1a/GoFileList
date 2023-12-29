package main

import "fmt"

func Palindrome() {
	var input string
	fmt.Scan(&input)

	var (
		inputRune = []rune(input)
		output    string
	)
	for i, k := 0, len(inputRune)-1; i < len(inputRune)/2; i++ {
		if inputRune[i] != inputRune[k] {
			fmt.Println("It is not a palindrome")
			return
		}
		output += string(inputRune[i])
		k--
	}
	fmt.Println(output)
}

func main() {
	Palindrome()
}
