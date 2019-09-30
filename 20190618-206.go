package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func reverseList(head *ListNode) *ListNode {
// 	var reversedList *ListNode = nil

// 	for head != nil {
// 		previous := head
// 		current := head.Next
// 		tmp := head.Next.Next

// 		current.Next = previous
// 		// tmp.Next = current
// 	}

// 	return head
// }

// func main() {
// 	var currentNode *ListNode = nil
// 	for index := 0; index < 4; index++ {
// 		currentNode = &ListNode{Val: index, Next: currentNode}
// 		fmt.Println(currentNode.Val)
// 	}

// 	loopNode := currentNode
// 	for loopNode.Next != nil {
// 		fmt.Println(loopNode.Val)
// 		loopNode = loopNode.Next
// 	}

// 	currentNode = reverseList(currentNode)

// 	loopNode = currentNode
// 	for loopNode.Next != nil {
// 		fmt.Println(loopNode.Val)
// 		loopNode = loopNode.Next
// 	}
// }
