package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func reversePrint(head *ListNode) []int {
// 	var stack []int
// 	node := head

// 	for node != nil {
// 		stack = append(stack, node.Val)
// 		node = node.Next
// 	}

// 	i, j := 0, len(stack)-1
// 	for i < j {
// 		stack[i], stack[j] = stack[j], stack[i]
// 	}
// 	return stack
// }

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	return append(reversePrint(head.Next), head.Val)
}
