package main

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestValue(root *TreeNode, target float64) int {
	left, right := root.Val, root.Val

	for root != nil {
		fmt.Println(left, right)
		if target < float64(root.Val) {
			right = root.Val
			root = root.Left
		} else if target > float64(root.Val) {
			left = root.Val
			root = root.Right
		} else {
			return root.Val
		}
	}
	lDelta, rDelta := math.Abs(float64(left)-target), math.Abs(float64(right)-target)
	if lDelta > rDelta {
		return right
	}

	return left
}

func main() {
	fmt.Println(1 - 1.0)
}
