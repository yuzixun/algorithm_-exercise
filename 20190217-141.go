package main

import "fmt"

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func hasCycle(head *ListNode) bool {
	pWalker := head
	pRunner := head

	for pWalker != nil && pRunner != nil && pWalker.Next != nil && pRunner.Next != nil {
		pWalker = pWalker.Next
		pRunner = pRunner.Next.Next

		if pWalker == pRunner {
			return true
		}
	}
	return false
}

func main() {
	var currentNode *ListNode = nil
	for index := 0; index < 4; index++ {
		currentNode = &ListNode{Val: index, Next: currentNode}
	}
	fmt.Println(hasCycle(currentNode))
}
