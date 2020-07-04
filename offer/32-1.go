package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	nodeArr := []*TreeNode{root}
	result := []int{}
	helperLevelOrder1(&nodeArr, 0)

	for i := 0; i < len(nodeArr); i++ {
		result = append(result, nodeArr[i].Val)
	}
	return result
}

func helperLevelOrder1(arrNode *[]*TreeNode, index int) {
	if len(*arrNode) == index {
		return
	}

	root := (*arrNode)[index]
	if root.Left != nil {
		(*arrNode) = append((*arrNode), root.Left)
		index++
		helperLevelOrder1(arrNode, index)
	}
	if root.Right != nil {
		(*arrNode) = append((*arrNode), root.Right)
		index++
		helperLevelOrder1(arrNode, index)
	}

}
