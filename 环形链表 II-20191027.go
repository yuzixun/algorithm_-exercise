package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	if fast != nil && fast.Next != nil {
		return nil
	}

	slow = head

	for {
		if slow == fast {
			return slow
		}
		slow = slow.Next
		fast = fast.Next
	}

	return nil
}
