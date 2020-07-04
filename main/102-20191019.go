package main

import . "../algorithm/treenode"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func bfs(node *TreeNode, level int, pResult *[][]int) {
	if node == nil {
		return
	}

	// result := *pResult
	size := len(*pResult)

	if level == size {
		// fmt.Println(size, pResult)
		(*pResult) = append((*pResult), []int{})
		size = len(*pResult)
	}
	// fmt.Println(*pResult, level)
	(*pResult)[level] = append((*pResult)[level], node.Val)

	bfs(node.Left, level+1, pResult)
	bfs(node.Right, level+1, pResult)
}

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	bfs(root, 0, &result)
	return result
}

// func main() {
// 	arr := []int{3, 9, 20, -1, -1, 15, 7}

// 	fmt.Println(levelOrder(generateTreeRoot(arr)))
// }
