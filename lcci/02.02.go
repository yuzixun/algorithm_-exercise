package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func kthToLast(head *ListNode, k int) int {
	count := 0
	current := head
	for current != nil {
		current = current.Next
		count++
	}

	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow.Val
}
