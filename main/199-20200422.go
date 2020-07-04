package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	result := []int{}

	list := []*TreeNode{root}
	nextList := []*TreeNode{}
	for len(list) != 0 {
		curRight := 0
		curLen := len(list)
		for i := 0; i < curLen; i++ {
			node := list[i]
			curRight = node.Val

			if node.Left != nil {
				nextList = append(nextList, node.Left)
			}
			if node.Right != nil {
				nextList = append(nextList, node.Right)
			}
		}
		result = append(result, curRight)
		list, nextList = nextList, []*TreeNode{}
	}

	return result
}
