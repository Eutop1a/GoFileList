package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
)

func main() {
	var (
		method    string
		plaintext string
		textBytes []byte
	)
	flag.StringVar(&method, "method", "1", "encoding type 1: sha256, 2: sha384, 3: sha512")
	flag.Parse()
	validateMethod(method)

	fmt.Println("Please Enter a string")
	fmt.Scan(&plaintext)
	textBytes = []byte(plaintext)
	if method == "1" {
		fmt.Printf("sha256: %x\n", sha256.Sum256(textBytes))
	} else if method == "2" {
		fmt.Printf("sha384: %x\n", sha512.Sum384(textBytes))
	} else {
		fmt.Printf("sha512: %x\n", sha512.Sum512(textBytes))
	}
}

func validateMethod(method string) {
	if method != "1" && method != "2" && method != "3" {
		log.Fatalf("method %s is not supported", method)
	}
}
