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
	value := 0
	helper(root, &value)
	return value
}

func helper(node *TreeNode, cur *int) int {
	if node == nil {
		return 0
	}

	left, right := 0, 0
	if node.Left != nil {
		left = helper(node.Left, cur)
	}

	if node.Right != nil {
		right = helper(node.Right, cur)
	}

	oneSide := max(0, max(left, right)) + root.Val
	twoSide := max(left, 0) + max(right, 0) + root.Val

	(*cur) = max((*cur), max(oneSide, twoSide))
	return oneSide
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
