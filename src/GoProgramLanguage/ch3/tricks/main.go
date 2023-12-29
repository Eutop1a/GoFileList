package main

import "fmt"

func main() {
	test := 0666
	fmt.Printf("%d, %[1]o, %#[1]o\n", test)

	ascii := 'a'
	unicode := '国'
	fmt.Printf("%d, %[1]c, %[1]q\n", ascii)

	fmt.Printf("%d, %[1]c, %[1]q\n", unicode)
}
