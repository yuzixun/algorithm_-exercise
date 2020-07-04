package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	count := 0
	cur := head
	for cur != nil {
		count++
		cur = cur.Next
	}
	cur = head
	count = count/2 + 1
	for i := 0; i < count; i++ {
		cur = cur.Next
	}
	return cur
}
