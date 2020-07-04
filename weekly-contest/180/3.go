package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func balanceBST(root *TreeNode) *TreeNode {
	var arr []int
	dfs(root, &arr)

	return rebuild(0, len(arr)-1, arr)
}

func dfs(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, arr)
	(*arr) = append((*arr), root.Val)
	dfs(root.Right, arr)
}

func rebuild(left, right int, arr []int) *TreeNode {
	if left > right {
		return &TreeNode{}
	}

	cur := left + (right-left)/2
	node := &TreeNode{Val: arr[cur]}
	node.Left = rebuild(left, cur-1, arr)
	node.Right = rebuild(cur+1, right, arr)

	return node
}
