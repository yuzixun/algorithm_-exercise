package main

// func reverseList(head *ListNode) *ListNode {
// 	var newHead *ListNode = nil

// 	for head != nil {
// 		tmpHead := head.Next
// 		head.Next = newHead
// 		newHead = head
// 		head = tmpHead
// 	}

// 	return newHead
// }

// // func main() {
// // 	var currentNode *ListNode = nil
// // 	for index := 0; index < 4; index++ {
// // 		currentNode = &ListNode{Val: index, Next: currentNode}
// // 		// fmt.Println(currentNode.Val)
// // 	}[]int{}

// // 	loopNode := currentNode
// // 	for loopNode.Next != nil {
// // 		// fmt.Println(loopNode.Val)
// // 		loopNode = loopNode.Next
// // 	}

// // 	currentNode = reverseList(currentNode)

// // 	loopNode = currentNode
// // 	for loopNode.Next != nil {
// // 		fmt.Println(loopNode.Val)
// // 		loopNode = loopNode.Next
// // 	}
// // }
