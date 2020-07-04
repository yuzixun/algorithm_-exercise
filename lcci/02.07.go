package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	reversedA, reversedB := false, false

	for a != b {
		if a.Next == nil {
			if reversedA {
				break
			}
			a = headB
			reversedA = true
		}

		if b.Next == nil {
			if reversedB {
				break
			}
			b = headA
			reversedB = true
		}

		a = a.Next
		b = b.Next
	}

	return nil
}
