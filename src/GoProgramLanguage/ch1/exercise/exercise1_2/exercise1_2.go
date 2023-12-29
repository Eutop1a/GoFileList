package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	line2file := make(map[string]string)

	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Println("Please input one or more file names")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, line2file)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, line2file[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, line2file map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		line2file[input.Text()] = f.Name()
	}
}
