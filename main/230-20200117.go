package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func kthSmallest(root *TreeNode, k int) int {
// 	var arr []int
// 	middleOrder(root, &arr)
// 	return arr[k]
// }

// func middleOrder(root *TreeNode, arr *[]int) {
// 	if root == nil {
// 		return
// 	}
// 	middleOrder(root.Left, arr)
// 	(*arr) = append((*arr), root.Val)
// 	middleOrder(root.Right, arr)
// }

func kthSmallest(root *TreeNode, k int) int {
	var (
		res  *TreeNode
		find func(root *TreeNode)
	)

	find = func(node *TreeNode) {
		if node.Left != nil {
			find(node.Left)
		}
		k--
		if k == 0 {
			res = node
			return
		}
		if node.Right != nil {
			find(node.Right)
		}

	}

	find(root)
	return res.Val
}
