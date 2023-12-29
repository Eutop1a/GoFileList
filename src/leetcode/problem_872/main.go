package main

import (
	"fmt"
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var list []int
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			list = append(list, root.Val)
			return
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root1)
	vals := append([]int(nil), list...)
	list = []int{}
	dfs(root2)
	return reflect.DeepEqual(list, vals)
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
	fmt.Println(leafSimilar(root, root))
	// fmt.Println(maxDepth(root))

}
