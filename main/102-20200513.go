package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}
	nextQueue := []*TreeNode{}
	for len(queue) != 0 {
		cur := []int{}

		for _, node := range queue {
			cur = append(cur, node.Val)

			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}

			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}

		queue = nextQueue
		nextQueue = []*TreeNode{}
		result = append(result, cur)
	}

	return result
}
