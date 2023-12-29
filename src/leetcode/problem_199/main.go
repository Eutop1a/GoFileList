package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func rightSideView(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		ans = append(ans, queue[len(queue)-1].Val)
		help := queue
		queue = nil
		for _, v := range help {
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
	}
	return ans
}

func rightSideView2(root *TreeNode) []int {
	var ans []int
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth >= len(ans) {
			ans = append(ans, node.Val)
		}
		dfs(node.Right, depth+1) // 优先从右边开始遍历
		dfs(node.Left, depth+1)
	}
	dfs(root, 0)
	return ans
}

func main() {
	root := newTreeNode(3)
	root.Left = newTreeNode(5)
	root.Right = newTreeNode(1)
	root.Left.Left = newTreeNode(6)
	root.Left.Right = newTreeNode(2)
	root.Left.Right.Left = newTreeNode(7)
	root.Left.Right.Right = newTreeNode(4)
	root.Right.Left = newTreeNode(9)
	root.Right.Right = newTreeNode(8)
	fmt.Println(rightSideView2(root))

}
