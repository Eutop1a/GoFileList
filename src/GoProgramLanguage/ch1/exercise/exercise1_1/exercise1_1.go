package main

import (
	"fmt"
	"os"
)

func main() {
	for index, value := range os.Args {
		if index == 0 {
			fmt.Println("os.Args[0] =", value)
		} else {
			fmt.Println(index, "=", value)
		}

	}
}
