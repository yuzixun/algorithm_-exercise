package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
	prev, cur := head, head.Next

	if prev.Val == val {
		return cur
	}

	for cur != nil {
		if cur.Val == val {
			prev.Next = cur.Next
			return head
		}

		prev = cur
		cur = cur.Next
	}

	return head
}
