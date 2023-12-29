package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

}
