package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	result := -1 << 32
	maxPath(root, &result)
	return result
}

func maxPath(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	leftVal, rightVal := 0, 0
	if root.Left != nil {
		leftVal = maxPath(root.Left, result)
	}

	if root.Right != nil {
		rightVal = maxPath(root.Right, result)
	}

	oneSide := max(0, max(leftVal, rightVal)) + root.Val
	twoSide := max(leftVal, 0) + max(rightVal, 0) + root.Val

	*result = max(*result, max(oneSide, twoSide))
	return oneSide
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
