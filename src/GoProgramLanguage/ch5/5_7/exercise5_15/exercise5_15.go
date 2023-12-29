package main

import "fmt"

func Max(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	maxVal := -(2<<32 - 1)
	for _, val := range vals {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func Min(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	minVal := 2<<32 - 1
	for _, val := range vals {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

func main() {
	fmt.Println(Max(1, 2, 3, 4, 5, 6, -100))
	fmt.Println(Min(1, 2, 3, 4, 5, 6, -100))
	fmt.Println(Max())
	fmt.Println(Min())
}
