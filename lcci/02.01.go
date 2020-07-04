package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cache := make(map[int]int, 0)
	prev, current := head, head
	for current != nil {
		_, ok := cache[current.Val]
		if ok {
			continue
		}
		cache[current.Val]++
		prev, current = current, current.Next
	}

	return head
}
