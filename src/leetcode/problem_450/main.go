package main

import "fmt"

var fatherNode *TreeNode
var depth int

func find(node *TreeNode, key int) *TreeNode {
	depth = 0
	fatherNode = nil
	for node != nil {
		if depth == 0 { // if key == root.Val
			fatherNode = node
		}
		if key == node.Val {
			return node
		}
		fatherNode = node
		depth++
		if key > node.Val {
			node = node.Right
		} else {
			node = node.Left
		}
	}
	return nil
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	keyNode := find(root, key)
	if keyNode == nil {
		return root
	}
	if fatherNode.Val > keyNode.Val {
		if keyNode.Left != nil && keyNode.Right != nil {
			fatherNode.Left = keyNode.Right
			keyNode.Right.Left = keyNode.Left
		} else if keyNode.Left != nil && keyNode.Right == nil {
			fatherNode.Left = keyNode.Left
		} else if keyNode.Left == nil && keyNode.Right != nil {
			fatherNode.Left = keyNode.Right
		} else {
			fatherNode.Left = nil
		}
	} else if fatherNode.Val < keyNode.Val {
		if keyNode.Left != nil && keyNode.Right != nil {
			fatherNode.Right = keyNode.Right
			keyNode.Right.Left = keyNode.Left
		} else if keyNode.Left != nil && keyNode.Right == nil {
			fatherNode.Right = keyNode.Left
		} else if keyNode.Left == nil && keyNode.Right != nil {
			fatherNode.Right = keyNode.Right
		} else {
			fatherNode.Right = nil
		}
	} else { // fatherNode.Val == keyNode.Val
		if keyNode.Left != nil && keyNode.Right != nil {
			if keyNode.Right.Left == nil {
				keyNode.Right.Left = keyNode.Left
				return keyNode.Right
			} else {
				keyNode.Right.Left.Left = keyNode.Left
				keyNode.Right.Left.Right = keyNode.Right
				temp := keyNode.Right.Left
				keyNode.Right.Left = nil
				return temp
			}
		} else if keyNode.Left != nil && keyNode.Right == nil {
			return keyNode.Left
		} else if keyNode.Left == nil && keyNode.Right != nil {
			return keyNode.Right
		} else {
			return nil
		}

	}
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func main() {
	root := newTreeNode(50)
	root.Left = newTreeNode(30)
	root.Right = newTreeNode(70)
	root.Left.Right = newTreeNode(40)
	root.Right.Left = newTreeNode(60)
	root.Right.Right = newTreeNode(80)
	test := deleteNode(root, 50)
	fmt.Println(test)
	fmt.Println(fatherNode.Val)

	root1 := newTreeNode(0)
	fmt.Println(deleteNode(root1, 0))

}
