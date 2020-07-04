package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1, stack2 := []int{}, []int{}

	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}
	var node *ListNode
	carry, cur := 0, 0
	len1, len2 := len(stack1), len(stack2)
	for len1 > 0 && len2 > 0 {
		cur := stack1[len1] + stack2[len2] + carry
		carry = cur % 10
		cur = cur / 10
		node = &ListNode{Val: cur, Next: node}
	}

	for len1 > 0 {
		cur := stack1[len1] + carry
		carry = cur % 10
		cur = cur / 10
		node = &ListNode{Val: cur, Next: node}
	}
	for len2 > 0 {
		cur := stack2[len2] + carry
		carry = cur % 10
		cur = cur / 10
		node = &ListNode{Val: cur, Next: node}
	}
	if carry != 0 {
		node = &ListNode{Val: carry, Next: node}
	}
	return node
}
