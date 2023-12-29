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

func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if val == root.Val {
			return root
		}
		if val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return nil
}

func searchBST2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if val == root.Val {
		return root
	}
	if val < root.Val {
		return searchBST2(root.Left, val)
	}
	return searchBST2(root.Right, val)
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
	fmt.Println(searchBST(root, 3))
	fmt.Println(searchBST2(root, 3))

}
