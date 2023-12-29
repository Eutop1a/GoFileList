package main

import "fmt"

func maxLevelSum(root *TreeNode) int {
	var res int
	var q []int
	var bfs func(*TreeNode)
	bfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		q = append(q, node.Val)
		var tmp int
		for _, v := range q {
			tmp += v
		}
		res = max(res, tmp)
		bfs(node.Left)
		bfs(node.Right)
	}
	bfs(root)
	return res
}

func maxLevelSum2(root *TreeNode) int {
	ans, maxSum := 1, root.Val
	q := []*TreeNode{root}
	for level := 1; len(q) > 0; level++ {
		tmp := q
		q = nil
		sum := 0
		for _, node := range tmp {
			sum += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if sum > maxSum {
			ans, maxSum = level, sum
		}
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

// 深度优先搜索
func dfs(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, node.Val)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return result
}

// 广度优先搜索
func bfs(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		result = append(result, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}

// 深度优先搜索（递归版本）
func dfsRecursive(root *TreeNode) []int {
	var result []int
	dfsHelper(root, &result)
	return result
}

func dfsHelper(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	*result = append(*result, node.Val)

	dfsHelper(node.Left, result)
	dfsHelper(node.Right, result)
}

// 广度优先搜索（递归版本）
func bfsRecursive(root *TreeNode) []int {
	var result []int
	var queue []*TreeNode
	queue = append(queue, root)
	bfsRecursiveHelper(queue, &result)
	return result
}

func bfsRecursiveHelper(queue []*TreeNode, result *[]int) {
	if len(queue) == 0 {
		return
	}

	node := queue[0]
	queue = queue[1:]

	*result = append(*result, node.Val)

	if node.Left != nil {
		queue = append(queue, node.Left)
	}
	if node.Right != nil {
		queue = append(queue, node.Right)
	}

	bfsRecursiveHelper(queue, result)
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
	fmt.Println(dfs(root))
	fmt.Println(dfsRecursive(root))
	fmt.Println(bfs(root))
	fmt.Println(bfsRecursive(root))
}

var q []int

func BFSPrintBinTree(node *TreeNode) {
	if node == nil {
		return
	}
	q = append(q, node.Val)
	BFSPrintBinTree(node.Left)
	BFSPrintBinTree(node.Right)
}
