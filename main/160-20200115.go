package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	countA, countB := 0, 0
// 	tmpA, tmpB := headA, headB

// 	for tmpA != nil {
// 		tmpA = tmpA.Next
// 		countA++
// 	}
// 	for tmpB != nil {
// 		tmpB = tmpB.Next
// 		countB++
// 	}

// 	tmpA, tmpB = headA, headB
// 	for countA > countB {
// 		tmpA = tmpA.Next
// 		countA--
// 	}
// 	for countB > countA {
// 		tmpB = tmpB.Next
// 		countB--
// 	}

// 	for tmpA != nil && tmpB != nil {
// 		if tmpA == tmpB {
// 			return tmpA
// 		}
// 		tmpA = tmpA.Next
// 		tmpB = tmpB.Next
// 	}
// 	return nil
// }

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	tmpA, tmpB := headA, headB

	for tmpA != tmpB {

		if tmpA == nil {
			tmpA = headB
		} else {
			tmpA = tmpA.Next
		}

		if tmpB == nil {
			tmpB == headA
		} else {
			tmpB = tmpB.Next
		}

	}

	return tmpA
}
