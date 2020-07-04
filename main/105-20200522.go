package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root, preorder := preorder[0], preorder[1:]

	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == root {
			break
		}
	}

	node := &TreeNode{Val: root}
	node.Left = buildTree(preorder[:i], inorder[:i])
	node.Right = buildTree(preorder[i:], inorder[i+1:])

	return node
}
