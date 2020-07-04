package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0

	head := &ListNode{}
	current := head
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry
		current.Next = &ListNode{Val: sum % 10}

		carry = sum / 10

		current = current.Next
		l1, l2 = l1.Next, l2.Next
	}

	for l1 != nil {
		sum := l1.Val + carry
		current.Next = &ListNode{Val: sum % 10}

		carry = sum / 10

		current = current.Next
		l1 = l1.Next
	}

	for l2 != nil {
		sum := l2.Val + carry
		current.Next = &ListNode{Val: sum % 10}

		carry = sum / 10

		current = current.Next
		l2 = l2.Next
	}
	return head.Next
}
