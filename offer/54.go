package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargest(root *TreeNode, k int) int {
	arr := []int{}

	dfs(root, &arr, k)
	return arr[k-1]
}

func dfs(root *TreeNode, arr *[]int, k int) {
	if len(*arr) >= k {
		return
	}

	if root.Right != nil {
		dfs(root.Right, arr, k)
	}
	(*arr) = append((*arr), root.Val)

	if root.Left != nil {
		dfs(root.Left, arr, k)
	}
}
