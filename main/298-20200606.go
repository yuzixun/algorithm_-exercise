package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestConsecutive(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftRes, rightRes := 0, 0
	if root.Left != nil {
		if root.Left.Val == root.Val+1 {
			leftRes = longestConsecutive(root.Left) + 1
		} else {
			leftRes = longestConsecutive(root.Left)
		}
	}

	if root.Right != nil {
		if root.Right.Val == root.Val+1 {
			rightRes = longestConsecutive(root.Right) + 1
		} else {
			rightRes = longestConsecutive(root.Right)
		}
	}

	return max(leftRes, rightRes)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestConsecutive(root *TreeNode) int {
	return dfs(nil, root, 0)
}

func dfs(parent, child *TreeNode, size int) int {
	if child == nil {
		return size
	}

	if parent != nil && parent.Val+1 == child.Val {
		size++
	} else {
		size = 1
	}

	return max(size, max(dfs(child, child.Left, size), dfs(child, child.Right, size)))
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {

}
