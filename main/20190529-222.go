package main

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

func countNodes(root *TreeNode) int {
	return counter(0, root)
}

func counter(count int, root *TreeNode) int {
	if root == nil {
		return count
	}
	if root.Left != nil {
		count = counter(count, root.Left)
	}
	if root.Right != nil {
		count = counter(count, root.Right)
	}

	return count + 1
}

// func main() {

// }
