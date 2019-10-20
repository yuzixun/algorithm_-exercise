package treenode

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var NULL = math.Pow(2, 32) - 1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GenerateTreeRoot(arr []int) *TreeNode {
	var node *TreeNode
	size := len(arr)

	root := &TreeNode{Val: arr[0]}
	queue := make([]*TreeNode, 1, size*2)
	queue[0] = root

	i := 1
	for i < size {
		node, queue = queue[0], queue[1:]

		if i < size && arr[i] != NULL {
			node.Left = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < size && arr[i] != NULL {
			node.Right = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
