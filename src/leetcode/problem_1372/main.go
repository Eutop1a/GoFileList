package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	res := 0
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, numLeft int, numRight int) {
		fmt.Println(node.Val)
		if node.Left != nil {
			numRight++
			if numRight > res {
				res = numRight
			}
			dfs(node.Left, numRight, 0)
		}
		if node.Right != nil {
			numLeft++
			if numLeft > res {
				res = numLeft
			}
			dfs(node.Right, 0, numLeft)
		}
	}
	dfs(root, 0, 0)
	return res
}

/* true => left, false => right */

var maxAns int

func dfs(node *TreeNode, dir bool, len int) {
	maxAns = max(maxAns, len)
	if dir { // turn left
		if node.Left != nil {
			dfs(node.Left, false, len+1)
		}
		if node.Right != nil {
			dfs(node.Right, true, 1)
		}
	} else { // turn right
		if node.Right != nil {
			dfs(node.Right, true, len+1)
		}
		if node.Left != nil {
			dfs(node.Left, false, 1)
		}
	}

}

func longestZigZag2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	dfs(root, true, 0)
	dfs(root, false, 0)
	return maxAns
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
	fmt.Println(longestZigZag2(root))
}
