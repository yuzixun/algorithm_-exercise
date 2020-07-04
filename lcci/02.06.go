package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	left := head
	right := slow.Next

	for slow != nil {
		slow, slow.Next, right = slow.Next, right, slow
	}

	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next

	}
	return true
}
