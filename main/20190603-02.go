package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := new(ListNode)
	current := head
	for l1 != nil || l2 != nil || carry != 0 {
		v1, v2 := 0, 0
		if l1 != nil {
			l1, v1 = l1.Next, l1.Val
		}
		if l2 != nil {
			l2, v2 = l2.Next, l2.Val
		}

		currentVal := v1 + v2 + carry
		carry = currentVal / 10

		current.Next = &ListNode{Val: currentVal % 10, Next: nil}
		current = current.Next
	}
	return head.Next
}

// func main() {
// 	arr1 := []int{5}
// 	arr2 := []int{5}
// 	// (2 -> 4 -> 3) + (5 -> 6 -> 4)
// 	printListNode(addTwoNumbers(generateListNodeBy(arr1), generateListNodeBy(arr2)))
// }
