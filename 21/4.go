package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxSumBST(root *TreeNode) int {

	if root == nil {
		return 0
	}
	result := 0
	maxSumBSTHelper(root, 0, &result)
	return result
}

func maxSumBSTHelper(root *TreeNode, cur int, result *int) {
	if root.Left == nil && root.Right == nil {
		result = root.Val
		return
	}

	if root.Left != nil && root.Val > root.Left.Val {
		maxSumBSTHelper(root.Left, cur+root.Left.Val, result)
		result = max()
	}

	if root.Right != nil && root.Val < root.Right.Val {
		maxSumBSTHelper(root.Right, cur+root.Right.Val, result)
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
