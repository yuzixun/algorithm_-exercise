package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubPath(head *ListNode, root *TreeNode) bool {
	if helperIsSubPath(head, root) {
		return true
	}

	result := false
	if root.Left != nil {
		result = result || isSubPath(head, root.Left)
	}

	if root.Right != nil {
		result = result || isSubPath(head, root.Right)
	}

	return result
}

func helperIsSubPath(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}

	return head.Val == root.Val && helperIsSubPath(head.Next, root.Left) && helperIsSubPath(head.Next, root.Right)
}
