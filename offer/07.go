package main

/**
 * Definition for a binary tree node.
//  * type TreeNode struct {
//  *     Val int
//  *     Left *TreeNode
//  *     Right *TreeNode
//  * }
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return helper(preorder, inorder)
}

func helper(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return &TreeNode{}
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	pos := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			pos = i
			break
		}
	}
	root.Left = helper(preorder[1:pos+1], inorder[:pos])
	root.Right = helper(preorder[pos+1:], inorder[pos+1:])

	return root
}

// func main() {
// 	fmt.Println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))

// }
