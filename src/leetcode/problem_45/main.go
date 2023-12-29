package main

import "fmt"

func jump(nums []int) (step int) {
	n := len(nums)
	maxPos, end := 0, 0
	for i := 0; i < n-1; i++ {
		if maxPos >= i {
			maxPos = max(maxPos, i+nums[i])
		}
		if i == end {
			end = maxPos
			step++
		}
	}
	return
}

func main() {
	testArr := []int{2, 3, 1, 1, 4}
	tmp := jump(testArr)
	fmt.Println(tmp)
}
