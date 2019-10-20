package main

import (
	. "../algorithm/treenode"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)

	// 如果一边为空，
	// 计算深度时，需要去计算非空一边的深度
	if left == 0 || right == 0 {
		return left + right + 1
	}

	if left > right {
		return right + 1
	} else {
		return left + 1
	}
}
