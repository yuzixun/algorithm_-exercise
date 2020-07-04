package main

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	arrNode := []*TreeNode{root}
	result := [][]int{}
	temp := []int{}
	reverse := false
	for len(arrNode) != 0 {
		temp = []int{}
		if reverse {
			for i := len(arrNode) - 1; i >= 0; i-- {
				temp = append(temp, arrNode[i].Val)
			}
		} else {
			for i := 0; i < len(arrNode); i++ {
				temp = append(temp, arrNode[i].Val)
			}
		}

		result = append(result, temp)

		arrNode = helperLevelOrder(arrNode)

		reverse = !reverse
	}
	return result
}

func helperLevelOrder(arrNode []*TreeNode) []*TreeNode {
	result := []*TreeNode{}

	for i := 0; i < len(arrNode); i++ {
		if arrNode[i].Left != nil {
			result = append(result, arrNode[i].Left)
		}

		if arrNode[i].Right != nil {
			result = append(result, arrNode[i].Right)
		}
	}

	return result
}
