package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	var res [][]int
	var cur []int
	helperPathSum(root, sum, cur, &res)
	return res
}

func helperPathSum(root *TreeNode, sum int, cur []int, res *[][]int) {
	// fmt.Println(root, sum, cur, res)
	if root == nil {
		return
	}
	cur = append(cur, root.Val)
	sum -= root.Val
	if sum == 0 && root.Left == nil && root.Right == nil {
		x := make([]int, len(cur))
		copy(x, cur)
		*res = append(*res, x)
	} else {
		helperPathSum(root.Left, sum, cur, res)
		helperPathSum(root.Right, sum, cur, res)
	}
	// fmt.Println("before is ", cur)
	cur = cur[:len(cur)-1]
	// fmt.Println("after is ", cur)
}
