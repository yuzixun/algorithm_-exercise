package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	MIN, MAX := -1<<32, 1<<32

	return isValidBSTHelper(root, MIN, MAX)
}

func isValidBSTHelper(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	if node.Val > min && node.Val < max {
		return isValidBSTHelper(node.Left, min, node.Val) && isValidBSTHelper(node.Right, node.Val, max)
	}

	return false
}
