package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	bT := time.Now()

	s, sep := "", ""
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	cost := time.Now()
	fmt.Println("Run time: ", cost.Sub(bT))

	bT = time.Now()

	fmt.Println(strings.Join(os.Args, " "))

	cost = time.Now()

	fmt.Println("Run time: ", cost.Sub(bT))
}
