package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 寻找以root为根节点，是否有路径总和为targetSum的路径
func rootSum(root *TreeNode, targetSum int) (res int) {
	if root == nil {
		return
	}
	val := root.Val
	if val == targetSum {
		res++
	}
	res += rootSum(root.Left, targetSum-val)
	res += rootSum(root.Right, targetSum-val)
	return
}

// 用于遍历整个二叉树，把每一个节点都当一次根节点
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := rootSum(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}

func newTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func pathSum2(root *TreeNode, targetSum int) (ans int) {
	perSum := map[int]int{0: 1}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, curr int) {
		if node == nil {
			return
		}
		curr += node.Val
		ans += perSum[curr-targetSum]
		perSum[curr]++
		dfs(node.Left, curr)
		dfs(node.Right, curr)
		perSum[curr]--
		return
	}
	dfs(root, 0)
	return
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
	fmt.Println(pathSum2(root, 8))
	// fmt.Println(maxDepth(root))

}
