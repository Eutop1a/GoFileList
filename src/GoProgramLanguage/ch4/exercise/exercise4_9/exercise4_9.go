package main

import (
	"bufio"
	"fmt"
	"os"
)

func wordFreq() {
	hash := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		hash[input.Text()]++
	}
	fmt.Printf("%#v\n", hash)
}

func main() {
	wordFreq()
}
