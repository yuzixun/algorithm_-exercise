package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	pre, tail := dummy, dummy

	for tail.Next != nil {

		for i := 0; i < k && tail != nil; i++ {
			tail = tail.Next
		}
		if tail == nil {
			break
		}
		next := tail.Next
		tail.Next = nil

		cur := pre.Next

		pre.Next = reverseList(cur)

		cur.Next = next

		pre = cur
		tail = cur

	}

	return dummy.Next
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}

func main() {

}
