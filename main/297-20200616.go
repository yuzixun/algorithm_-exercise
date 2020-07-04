package main

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {

}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}

	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

func dfs(vals *[]string) *TreeNode {
	val := (*vals)[0]
	(*vals) = (*vals)[1:]

	if val == "#" {
		return nil
	}

	nVal, _ := strconv.Atoi(val)
	node := &TreeNode{Val: nVal}
	node.Left = dfs(vals)
	node.Right = dfs(vals)
	return node
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	return dfs(&vals)
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
func main() {

}
