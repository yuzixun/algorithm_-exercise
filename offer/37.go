package main

import (
	"fmt"
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
func Codec(root *TreeNode) string {
	if root == nil {
		return "[]"
	}
	arr := []*TreeNode{root}

	result := "["

	for len(arr) != 0 {
		cur := arr[0]
		arr = arr[1:]

		if cur != nil {
			result += fmt.Sprintf("%v,", cur.Val)
			arr = append(arr, cur.Left)
			arr = append(arr, cur.Right)
		} else {
			result += "null,"
		}

	}
	result = result[:len(result)-1]
	result += "]"
	return result
}

func deCodec(data string) *TreeNode {
	if data == "" || data == "[]" {
		return nil
	}

	data = data[1 : len(data)-1]
	arr := strings.Split(data, ",")
	index := 0

	root := generateTreeNode(arr[index])
	index++

	node := root

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		node, queue = queue[0], queue[1:]
		node.Left = generateTreeNode(arr[index])
		index++
		node.Right = generateTreeNode(arr[index])
		index++

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return root
}

func generateTreeNode(value string) *TreeNode {
	if value == "null" {
		return nil
	}

	val, err := strconv.Atoi(value)
	if err != nil {
		return nil
	}

	return &TreeNode{Val: val}
}
