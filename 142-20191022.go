package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func detectCycle(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	arr := []*ListNode{}
// 	pointer := head
// 	for pointer != nil {
// 		for _, item := range arr {
// 			if item == pointer {
// 				return item
// 			}
// 		}
// 		arr = append(arr, pointer)
// 		pointer = pointer.Next
// 	}
// 	return nil
// }

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			break
		}
	}

	if fast.Next == nil || fast.Next.Next == nil {
		return nil
	}

	fast = head
	for fast != slow {
		fast, slow = fast.Next, slow.Next
	}

	return fast
}
