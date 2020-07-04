package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	size := len(lists)
	if size == 0 {
		return nil
	}
	if size == 1 {
		return lists[0]
	}
	if size == 2 {
		merge2Lists(lists[0], lists[1])
	}

	list1, list2 := lists[:size/2], lists[size/2:]
	return merge2Lists(mergeKLists(list1), mergeKLists(list2))
}

func merge2Lists(list1, list2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			cur.Next = list2
			list2 = list2.Next
		} else {
			cur.Next = list1
			list1 = list1.Next
		}
		cur = cur.Next
	}

	for list1 != nil {
		cur.Next = list1
		list1 = list1.Next
		cur = cur.Next
	}

	for list2 != nil {
		cur.Next = list2
		list2 = list2.Next
		cur = cur.Next
	}

	return head.Next
}

func main() {

}
