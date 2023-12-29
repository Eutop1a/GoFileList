package main

import (
	"fmt"
	"math"
	"sort"
)

func removeDuplicates0(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	fast, slow := 1, 0
	times := 1
	for ; fast < n; fast++ {
		if times < 2 && nums[fast] == nums[slow] {
			times++
		} else if times >= 2 && nums[fast] == nums[slow] {
			nums[fast] = math.MaxInt
		} else {
			times = 1
		}
		slow++
	}
	sort.Ints(nums)
	return slow
}
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	fast, slow := 2, 2
	for fast < n {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}

func main() {
	arr := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates(arr))
	fmt.Println(arr)
}
