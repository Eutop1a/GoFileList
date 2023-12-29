package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 使用递归
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// DFS 使用非递归
func DFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var stack = []*TreeNode{root}
	var depth = 0 // 初始化深度为0

	for len(stack) > 0 {
		size := len(stack) // 记录当前层的节点数
		depth++            // 深度加一，表示进入下一层

		// 遍历当前层的所有节点
		for i := 0; i < size; i++ {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		}
	}

	return depth
}

func newTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func main() {
	root := newTreeNode(1)
	root.Left = newTreeNode(2)
	root.Right = newTreeNode(3)
	root.Left.Left = newTreeNode(4)
	root.Right.Left = newTreeNode(5)
	// fmt.Println(maxDepth(root))
	fmt.Println(DFS(root))
}
