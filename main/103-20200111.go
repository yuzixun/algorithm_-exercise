package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	var arr []*TreeNode
	nextArr := []*TreeNode{root}
	leftToRight := false

	for len(nextArr) > 0 {
		var currentRes []int

		arr = nextArr
		nextArr = []*TreeNode{}
		leftToRight = !leftToRight

		for _, node := range arr {
			if node.Left != nil {
				nextArr = append(nextArr, node.Left)
			}

			if node.Right != nil {
				nextArr = append(nextArr, node.Right)
			}

			currentRes = append(currentRes, node.Val)
		}

		if !leftToRight {
			for i := len(currentRes)/2 - 1; i >= 0; i-- {
				opp := len(currentRes) - 1 - i
				currentRes[i], currentRes[opp] = currentRes[opp], currentRes[i]
			}
		}

		result = append(result, currentRes)
	}

	fmt.Println(result)
	return result
}

// func main() {

// }
