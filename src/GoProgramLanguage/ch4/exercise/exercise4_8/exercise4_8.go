package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	hash := make(map[string]int, 4)

	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		t := whichType(r)
		hash[t]++
	}
	fmt.Println(hash)
}

func whichType(r rune) string {
	if unicode.IsNumber(r) {
		return "number"
	}
	if unicode.Is(unicode.Han, r) {
		return "Chinese"
	}
	if r >= 64 && r <= 122 {
		return "English"
	}
	return "other"
}
