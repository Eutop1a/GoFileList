package main

import (
	"fmt"
	"sort"
	"time"
)

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func majorityElement2(nums []int) int {
	n := len(nums)
	recorder := map[int]int{}
	for i := 0; i < n; i++ {
		recorder[nums[i]]++
		if recorder[nums[i]] > n/2 {
			return nums[i]
		}
	}
	return 0
}

func majorityElement3(nums []int) int {
	n := len(nums)
	recorder := map[int]int{}
	for i := 0; i < n; i++ {
		recorder[nums[i]]++
	}
	for k, v := range recorder {
		if v > n/2 {
			return k
		}
	}
	return 0
}

func main() {
	// 生成一个包含大量数据的切片
	arr := []int{2, 2, 1, 1, 1, 2, 2, 2, 1, 1, 1, 2, 2, 2, 1, 1, 1, 2, 2, 2, 1, 1, 1, 2, 2, 2, 1, 1, 1, 2, 2, 2, 1, 1, 1, 2, 2, 2, 1, 1, 1, 2, 2, 2, 1, 1, 1, 2, 2, 2, 1, 1, 1, 2, 2, 2, 1, 1, 1, 2, 2, 2, 1, 1, 1, 2, 2, 2, 1, 1, 1, 2, 2, 2, 1, 1, 1, 2, 2, 2, 1, 1, 1, 2, 2, 2, 1, 1, 1, 2, 2}

	// 多次调用函数以增加执行时间
	numIterations := 10000000
	start := time.Now()
	for i := 0; i < numIterations; i++ {
		tmp := majorityElement3(arr)
		_ = tmp
	}
	elapsed := time.Since(start).Milliseconds()
	fmt.Printf("time cost = %v ms\n", elapsed)

	// 多次调用函数以增加执行时间
	start = time.Now()
	for i := 0; i < numIterations; i++ {
		tmp := majorityElement2(arr)
		_ = tmp
	}
	elapsed = time.Since(start).Milliseconds()
	fmt.Printf("time cost = %v ms\n", elapsed)
}
