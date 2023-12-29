package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X")) // [32]byte
	count := 0
	for i := 0; i < 32; i++ {
		if c1[byte(1<<(i*8))] != c2[byte(1<<(i*8))] {
			count++
		}
	}
	fmt.Println("different bits number is :", count)
}
