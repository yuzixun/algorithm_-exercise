package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	count := 0
	slow, fast := head, head

	for fast.Next != nil {
		fast = fast.Next
		count++
	}

	rk := k % count
	fast = head
	count = 0

	for fast.Next != nil {
		fast = fast.Next
		if count >= rk {
			slow = slow.Next
		}
		count++
	}

	newHead := slow.Next
	slow.Next = nil
	fast.Next = head

	return newHead
}

// func main() {

// }
