package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {

	visited := map[*Node]*Node{}

	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		if node == nil {
			return nil
		}

		if v, ok := visited[node]; ok {
			return v
		}

		dup := &Node{Val: node.Val}
		visited[node] = dup

		for _, neighbor := range node.Neighbors {
			dup.Neighbors = append(dup.Neighbors, dfs(neighbor))
		}

		return dup
	}

	return dfs(node)
}
