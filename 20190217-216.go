package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode = nil

	for head != nil {
		tmpHead := head.Next
		head.Next = newHead
		newHead = head
		head = tmpHead
	}

	return newHead
}

func main() {
	var currentNode *ListNode = nil
	for index := 0; index < 4; index++ {
		currentNode = &ListNode{Val: index, Next: currentNode}
		// fmt.Println(currentNode.Val)
	}

	loopNode := currentNode
	for loopNode.Next != nil {
		// fmt.Println(loopNode.Val)
		loopNode = loopNode.Next
	}

	currentNode = reverseList(currentNode)

	loopNode = currentNode
	for loopNode.Next != nil {
		fmt.Println(loopNode.Val)
		loopNode = loopNode.Next
	}
}
