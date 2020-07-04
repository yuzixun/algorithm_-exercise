package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func generateListNodeBy(array []int) *ListNode {
	var before *ListNode
	var first *ListNode

	for index, item := range array {
		current := &ListNode{Val: item}

		if index == 0 {
			first = current
		} else {
			before.Next = current
		}

		before = current
	}

	return first
}

func printListNode(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

}
