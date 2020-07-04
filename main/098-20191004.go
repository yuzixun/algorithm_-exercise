package main

// import . "../algorithm/treenode"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func isValidBST(root *TreeNode) bool {
	MIN, MAX := -1<<63, 1<<63-1
	return helper(root, MIN, MAX)
}

func helper(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	return min < root.Val && max > root.Val && helper(root.Left, min, root.Val) && helper(root.Right, root.Val, max)
}
