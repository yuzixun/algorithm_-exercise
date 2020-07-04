package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	a, b := helper(root)
	return max(a, b)
}

func helper(root *TreeNode) (int, int) {
	var left1, left2, right1, right2 int

	if root.Left != nil {
		left1, left2 = helper(root.Left)
	}

	if root.Right != nil {
		right1, right2 = helper(root.Right)
	}

	leftVal := max(left1, left2)
	rightVal := max(right1, right2)

	return root.Val + left2 + right2, leftVal + rightVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {

// }
