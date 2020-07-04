package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	result := []int{}
	edges := map[int][]int{}
	inDegree := map[int]int{}
	for _, prerequisite := range prerequisites {
		edges[prerequisite[1]] = append(edges[prerequisite[1]], prerequisite[0])
		inDegree[prerequisite[0]]++
	}

	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		result = append(result, cur)
		for _, item := range edges[cur] {
			inDegree[item]--
			if inDegree[item] == 0 {
				queue = append(queue, item)
			}
		}
	}

	if len(result) == numCourses {
		return result
	}
	return []int{}
}

func main() {
	fmt.Println(findOrder(2, [][]int{[]int{1, 0}}))
	fmt.Println(findOrder(4, [][]int{[]int{1, 0}, []int{2, 0}, []int{3, 1}, []int{3, 2}}))
}
