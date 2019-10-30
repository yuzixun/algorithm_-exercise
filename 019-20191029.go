package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// 已经走到最后一个节点，则说明要删除的是头结点
	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return head
}

// func main() {
// 	arr := []int{1, 2, 3, 4, 5}
// 	// arr := []int{1}

// 	node := &ListNode{Val: 0}
// 	head := node
// 	for _, val := range arr {
// 		node.Next = &ListNode{Val: val}
// 		node = node.Next
// 	}

// 	fmt.Println(removeNthFromEnd(head.Next, 2))
// }
