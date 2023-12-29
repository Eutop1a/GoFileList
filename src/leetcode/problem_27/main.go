package main

import (
	"fmt"
	"math"
	"sort"
)

func removeElement(nums []int, val int) int {
	n := len(nums)
	res := 0
	for i := 0; i < n; i++ {
		if nums[i] == val {
			res++
			nums[i] = math.MaxInt32
		}
	}
	sort.Ints(nums)
	return n - res
}

func removeElement2(nums []int, val int) int {
	n := len(nums)
	res := 0
	for i := 0; i < n; i++ {
		if nums[i] != val {
			nums[res] = nums[i]
			res++
		}
	}
	nums = nums[:res]
	return res
}

func main() {
	arr := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement2(arr, 2))
	fmt.Println(arr)
}
