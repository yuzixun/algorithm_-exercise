package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func isValidBST(root *TreeNode) bool {
// 	if root.Left == nil && root.Right == nil {
// 		return true
// 	}

// 	if root.Left == nil || root.Right == nil {
// 		return false
// 	}

// 	return helper(root, root.Left.Val, root.Right.Val)
// }

// func helper(root *TreeNode, min, max int) bool {
// 	if root.Val < min || root.Val > max {
// 		return false
// 	}

// 	return isValidBST(root.Left) && isValidBST(root.Right)
// }

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var pre *TreeNode

func isValidBST(root *TreeNode) bool {
	pre = nil
	sign := true

	helper(root, &sign)
	return sign
}

func helper(root *TreeNode, sign *bool) {
	if root != nil && (*sign) {
		helper(root.Left, sign)

		if pre != nil && root.Val >= pre.Val {
			*sign = false
		}

		pre = root
		helper(root.Right, sign)
	}
}

func main() {

}
