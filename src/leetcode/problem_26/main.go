package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	low := 0
	for fast := 0; fast < n-1; fast++ {
		if nums[fast] != nums[fast+1] {
			nums[low+1] = nums[fast+1]
			low++
		}
	}
	return low + 1
}

func removeDuplicates2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	low := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] { // 寻找新数组的元素 ，新数组就是不含有目标元素的数组
			nums[low] = nums[fast]
			low++ // 指向更新 新数组下标的位置
		}
	}
	return low
}

func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	tmp := removeDuplicates(arr)
	fmt.Println(arr, tmp)

}
