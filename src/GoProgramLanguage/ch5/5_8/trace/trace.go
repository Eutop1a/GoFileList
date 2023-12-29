package main

import (
	"fmt"
	"log"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// output: 	2023/12/03 20:06:46 enter bigSlowOperation
	//			2023/12/03 20:06:56 exit bigSlowOperation (10.0021308s)

	//defer trace("bigSlowOperation") //output:  2023/12/03 20:14:00 enter bigSlowOperation

	// ...lots of work…
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

func triple(x int) (result int) {
	defer func() {
		result += x
	}()
	return double(x)
}

func main() {
	//bigSlowOperation()
	fmt.Println(triple(4))
}
