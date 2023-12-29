package main

import (
	"fmt"
	"sort"
)

func merge0(nums1 []int, m int, nums2 []int, _ int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums := make([]int, m+n)
	i, j, k := 0, 0, 0
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			nums[k] = nums1[i]
			i++
			k++
		} else {
			nums[k] = nums2[j]
			j++
			k++
		}
	}
	// 将剩余的部分加入
	for i < m {
		nums[k] = nums1[i]
		i, k = i+1, k+1
	}
	for j < n {
		nums[k] = nums2[j]
		j, k = j+1, k+1
	}
	copy(nums1, nums)
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}

func main() {
	arr1 := []int{1, 2, 3, 0, 0, 0}
	arr2 := []int{2, 5, 6}
	merge2(arr1, 3, arr2, 3)
	fmt.Println(arr1)
}
