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
	result := [][]int{}

	bfs(root, 0, &result)

	return result
}

func bfs(node *TreeNode, level int, res *[][]int) {
	if node == nil {
		return
	}

	if level == len(*res) {
		(*res) = append((*res), []int{})
	}

	(*res)[level] = append((*res)[level], node.Val)

	bfs(node.Left, level+1, res)
	bfs(node.Right, level+1, res)

}
