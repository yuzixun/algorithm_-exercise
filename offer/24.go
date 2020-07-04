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

func reverseList(head *ListNode) *ListNode {
	prev, cur := head, head.Next

	for cur != nil {
		// fmt.Println(prev, cur, cur.Next)
		tmp := prev
		prev = cur
		cur = cur.Next
		cur.Next = prev
		// prev, cur, cur.Next = cur, cur.Next, prev
	}

	// fmt.Println(prev, prev.Next)
	return prev
}

// func main() {
// 	arr := []int{1, 2, 3, 4, 5}

// 	res := &ListNode{}
// 	head := res
// 	for i := 0; i < len(arr); i++ {
// 		head.Next = &ListNode{Val: arr[i]}
// 		head = head.Next
// 	}
// 	reverseList(res.Next)
// }
