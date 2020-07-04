package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func isBalanced(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}

// 	return math.Abs(depth(root.Left)-depth(root.Right)) <= 1.0 && isBalanced(root.Left) && isBalanced(root.Right)
// }

// func depth(root *TreeNode) float64 {
// 	if root == nil {
// 		return 0.0
// 	}

// 	return max(depth(root.Left), depth(root.Right)) + 1.0
// }

// func max(a, b float64) float64 {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	return helper(root) != -1
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left)
	if left == -1 {
		return -1
	}
	right := helper(root.Right)
	if right == -1 {
		return -1
	}

	delta := left - right
	if delta <= 1 && delta >= -1 {
		return max(left, right) + 1
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
