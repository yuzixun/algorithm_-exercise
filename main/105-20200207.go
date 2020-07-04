package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func buildTree(preorder []int, inorder []int) *TreeNode {
// 	return buildTreeHelper(preorder, inorder)
// }

// func buildTreeHelper(preorder, inorder []int) *TreeNode {
// 	if len(preorder) == 0 {
// 		return nil
// 	}

// 	val, preorder := preorder[0], preorder[1:]
// 	root := &TreeNode{Val: val}

// 	inPos := 0
// 	for i := 0; i < len(inorder); i++ {
// 		if inorder[i] == val {
// 			inPos = i
// 			break
// 		}
// 	}

// 	root.Left = buildTreeHelper(preorder[:inPos], inorder[:inPos])
// 	root.Right = buildTreeHelper(preorder[inPos:], inorder[inPos+1:])

// 	return root
// }

// func main() {
// 	a := []int{1, 2, 3}
// 	fmt.Println(a[:1])
// 	fmt.Println(a[2:])

// }
