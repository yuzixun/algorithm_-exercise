package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	a, cur := head, head
	for i := 0; i < k; i++ {
		cur = cur.Next
	}

	for cur != nil {
		cur = cur.Next
		a = a.Next
	}

	return a
}
