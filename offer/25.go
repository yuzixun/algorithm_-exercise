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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	head := result

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			result.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		} else {
			result.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		}
		// fmt.Println(result, l1, l2)
		result = result.Next
	}

	for l1 != nil {
		result.Next = &ListNode{Val: l1.Val}
		result = result.Next
		l1 = l1.Next
	}

	for l2 != nil {
		result.Next = &ListNode{Val: l2.Val}
		result = result.Next
		l2 = l2.Next
	}
	// fmt.Println(head.Next.Next)
	return head.Next
}

// func main() {

// 	a := buildListNode([]int{1, 2, 4})
// 	b := buildListNode([]int{1, 3, 4})
// 	mergeTwoLists(a, b)
// }

func buildListNode(arr []int) *ListNode {
	res := &ListNode{}
	head := res
	for i := 0; i < len(arr); i++ {
		head.Next = &ListNode{Val: arr[i]}
		head = head.Next
	}

	return res.Next
}
