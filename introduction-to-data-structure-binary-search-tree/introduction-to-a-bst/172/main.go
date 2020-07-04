package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	arr *[]int
	cur int
}

func Constructor(root *TreeNode) BSTIterator {
	bst := BSTIterator{arr: &[]int{}, cur: 0}
	bst.helper(root)
	return bst
}

func (this *BSTIterator) helper(root *TreeNode) {
	if root == nil {
		return
	}

	this.helper(root.Left)
	*this.arr = append(*this.arr, root.Val)
	this.helper(root.Right)
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	result := (*this.arr)[this.cur]
	this.cur++
	return result
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.cur < len(*this.arr)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
