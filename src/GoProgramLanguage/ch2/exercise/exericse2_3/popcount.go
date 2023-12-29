package main

import (
	"fmt"
	"time"
)

// pc[i] is the population count of i.
var pc [256]byte = func() (pc [256]byte) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return
}()

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCount2 returns the population count (number of set bits) of x.

func PopCount2(x uint64) int {
	count := 0
	for x != 0 {
		if x&1 == 1 {
			count++
		}
		x >>= 1
	}
	return count
}

func PopCount3(x uint64) int {
	var res int
	for i := 0; i < 8; i++ {
		res += int(pc[byte(x>>(i*8))])
	}
	return res
}

func main() {
	arr := make([]int, 10000000)
	for i := range arr {
		arr[i] = i + 1
	}
	fmt.Println(arr)
	start := time.Now()
	for _, v := range arr {
		PopCount(uint64(v))
	}
	fmt.Printf("pc cost: %d ms\n", time.Since(start).Milliseconds())

	start = time.Now()
	for _, v := range arr {
		PopCount2(uint64(v))
	}
	fmt.Printf("loop cost: %d ms\n", time.Since(start).Milliseconds())

	start = time.Now()
	for _, v := range arr {
		PopCount3(uint64(v))
	}
	fmt.Printf("loop2 cost: %d ms\n", time.Since(start).Milliseconds())
}
