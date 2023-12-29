package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	var DFS func(node *TreeNode, pathMax int) int
	DFS = func(node *TreeNode, pathMax int) int {
		if node == nil {
			return 0
		}
		res := 0
		if node.Val >= pathMax {
			res++
			pathMax = node.Val
		}
		res += DFS(node.Left, pathMax) + DFS(node.Right, pathMax)
		return res
	}
	return DFS(root, math.MinInt)
}

func dfs(root *TreeNode, pathMax int) int {
	if root == nil {
		return 0
	}
	res := 0
	if root.Val >= pathMax {
		res++
		pathMax = root.Val
	}
	res += dfs(root.Left, pathMax) + dfs(root.Right, pathMax)
	return res
}

func newTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
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
	fmt.Println(goodNodes(root))
}
