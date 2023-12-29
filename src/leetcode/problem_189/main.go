package main

import "fmt"

func rotate(nums []int, k int) {
	n := len(nums)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[(i+k)%n] = nums[i]

	}
	copy(nums, result)
}

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}

func rotate2(nums []int, k int) {
	k %= len(nums) // 防止k过大越界
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func rotate3(nums []int, k int) {
	n := len(nums)
	k %= n // 防止k过大越界
	arr := append(nums, nums[:n-k]...)
	copy(nums, arr[n-k:])
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	rotate3(arr, 1)
	fmt.Println(arr)
}
