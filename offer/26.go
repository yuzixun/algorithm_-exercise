package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func isSubStructure(A *TreeNode, B *TreeNode) bool {
// 	return helperSubStructure(A, B, B)
// }

// func helperSubStructure(A, B, subB *TreeNode) bool {
// 	fmt.Println(A, B, subB)
// 	if A == nil || subB == nil {
// 		return false
// 	}

// 	if A.Val == subB.Val {
// 		result := true
// 		if subB.Left != nil {
// 			result = result && helperSubStructure(A.Left, B, subB.Left)
// 		}
// 		if subB.Right != nil {
// 			result = result && helperSubStructure(A.Right, B, subB.Right)
// 		}
// 		// if subB.Left == nil && subB.Right == nil {
// 		// 	result
// 		// }
// 		return result
// 	}

// 	return helperSubStructure(A.Left, B, B) || helperSubStructure(A.Right, B, B)
// }

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	if helperSubStructure(A, B) {
		return true
	}
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func helperSubStructure(A, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}

	return A.Val == B.Val && helperSubStructure(A.Left, B.Left) && helperSubStructure(A.Right, B.Right)
}
