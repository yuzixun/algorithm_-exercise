package main

import "fmt"

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

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = nil
	l1 := sortList(head)
	l2 := sortList(slow)

	return merge(l1, l2)
}

func merge(l1, l2 *ListNode) *ListNode {
	c := &ListNode{Val: 0}
	p := c

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}

	if l1 != nil {
		p.Next = l1
	}

	if l2 != nil {
		p.Next = l2
	}

	return c.Next
}

func main() {
	vs := []int{-1, 5, 3, 4, 0}

	var cur, head *ListNode

	for _, v := range vs {
		fmt.Println(v, cur)
		if head != nil {
			cur.Next = &ListNode{Val: v}
			cur = cur.Next
		} else {
			head = &ListNode{Val: v}
			cur = head
		}
	}

	r := sortList(head)
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}
