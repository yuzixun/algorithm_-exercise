package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	var result []int
	process(root, &result)
	return result
}

func process(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		process(root.Left, result)
	}

	(*result) = append((*result), root.Val)

	if root.Right != nil {
		process(root.Right, result)
	}
}
