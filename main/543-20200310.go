package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var helper func(root *TreeNode, maxDiameter *int) int

	helper = func(root *TreeNode, maxDiameter *int) int {
		leftValue, rightValue := 0, 0
		if root.Left != nil {
			leftValue = helper(root.Left, maxDiameter)
		}
		if root.Right != nil {
			rightValue = helper(root.Right, maxDiameter)
		}

		if (*maxDiameter) < (leftValue + rightValue) {
			(*maxDiameter) = leftValue + rightValue
		}

		if leftValue > rightValue {
			return leftValue
		}
		return rightValue
	}

	maxDiameter := 0
	helper(root, maxDiameter)
	return maxDiameter
}

func main() {

}
