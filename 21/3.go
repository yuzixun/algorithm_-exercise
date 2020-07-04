package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxValue := 1
	longestZigZagHelper(root, false, 1, &maxValue)
	longestZigZagHelper(root, true, 1, &maxValue)

	fmt.Println(maxValue)
	return 0
}

func longestZigZagHelper(root *TreeNode, goLeft bool, cur int, maxValue *int) {

	if goLeft {
		if root.Left != nil {
			longestZigZagHelper(root.Left, false, cur+1, maxValue)
			longestZigZagHelper(root.Left, true, 1, maxValue)
		}
	} else {
		if root.Right != nil {
			longestZigZagHelper(root.Right, false, 1, maxValue)
			longestZigZagHelper(root.Right, true, cur+1, maxValue)
		}
	}

	(*maxValue) = max((*maxValue), cur)
}
