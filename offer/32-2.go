package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	arrNode := []*TreeNode{root}
	result := [][]int{}
	temp := []int{}
	for len(arrNode) != 0 {
		temp = []int{}
		for i := 0; i < len(arrNode); i++ {
			temp = append(temp, arrNode[i].Val)
		}
		result = append(result, temp)

		arrNode = helperLevelOrder2(arrNode)
	}
	return result
}

func helperLevelOrder2(arrNode []*TreeNode) []*TreeNode {
	result := []*TreeNode{}

	for i := 0; i < len(arrNode); i++ {
		if arrNode[i].Left != nil {
			result = append(result, arrNode[i].Left)
		}

		if arrNode[i].Right != nil {
			result = append(result, arrNode[i].Right)
		}
	}

	return result
}
