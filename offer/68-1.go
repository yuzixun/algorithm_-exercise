package main

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if (q.Val <= root.Val && root.Val <= q.Val) || (q.Val >= root.Val && root.Val >= q.Val) {
		return root
	}

	if q.Val > root.Val && p.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	if root.Val > q.Val && root.Val > p.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	return nil
}
