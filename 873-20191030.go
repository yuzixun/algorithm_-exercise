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

func middleNode(head *ListNode) *ListNode {

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// func main() {
// 	arr := []int{1, 2, 3, 4, 5, 6}

// 	node := &ListNode{Val: 0}
// 	head := node
// 	for _, val := range arr {
// 		node.Next = &ListNode{Val: val}
// 		node = node.Next
// 	}

// 	fmt.Println(middleNode(head.Next))
// }
