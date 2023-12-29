package main

import "fmt"

func canJump(nums []int) bool {
	n := len(nums)
	furthest := 0
	for i := 0; i < n; i++ {
		if i > furthest {
			return false
		}
		furthest = max(furthest, i+nums[i])
	}
	return true
}
func canJump2(nums []int) bool {
	furthest := 0
	n := len(nums) - 1
	for i := 0; i <= furthest; i++ { // 每次与覆盖值比较
		furthest = max(i+nums[i], furthest) //每走一步都将 furthest 更新为最大值
		if furthest >= n {
			return true
		}
	}
	return false
}

func main() {
	whether := canJump([]int{2, 3, 1, 1, 4})
	whether = canJump([]int{3, 2, 1, 0, 4})
	fmt.Println(whether)
}
