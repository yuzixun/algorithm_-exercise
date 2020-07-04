package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeDuplicateNodes(head *ListNode) *ListNode {
	usedMap := map[int]bool{}

	newHead := &ListNode{Val: 0}
	cur := newHead
	for head != nil {
		hasUsed := false

		fmt.Println(head.Val, hasUsed)
		if !usedMap[head.Val] {
			usedMap[head.Val] = true
			cur.Next = &ListNode{Val: head.Val}
			cur = cur.Next
		}
		head = head.Next
	}

	return newHead.Next
}
