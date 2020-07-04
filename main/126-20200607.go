package main

import "math"

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	ids := map[string]int{}
	result := [][]string{}

	for i, word := range wordList {
		ids[word] = i
	}

	if _, ok := ids[beginWord]; !ok {
		wordList = append(wordList, beginWord)
		ids[beginWord] = len(wordList) - 1
	}

	if _, ok := ids[endWord]; !ok {
		return result
	}

	n := len(wordList)
	edges := make([][]int, len(wordList))

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if transformCheck(wordList[i], wordList[j]) {
				edges[i] = append(edges[i], j)
				edges[j] = append(edges[j], i)
			}
		}
	}

	cost := make([]int, n)
	queue := [][]int{[]int{ids[beginWord]}}

	for i := 0; i < n; i++ {
		cost[i] = math.MaxInt32
	}
	cost[ids[beginWord]] = 0

	for i := 0; i < len(queue); i++ {
		cur := queue[i]
		last := cur[len(cur)-1]

		if last == ids[endWord] {
			tmp := []string{}
			for _, index := range cur {
				tmp = append(tmp, wordList[index])
			}
			result = append(result, tmp)
		} else {
			for _, to := range edges[last] {
				if cost[last]+1 <= cost[to] {
					cost[to] = cost[last] + 1
					tmp := make([]int, len(cur))
					copy(tmp, cur)
					tmp = append(tmp, to)
					queue = append(queue, tmp)
				}
			}
		}
	}
	return result
}

func transformCheck(from, to string) bool {
	for i := 0; i < len(from); i++ {
		if from[i] != to[i] {
			return from[i+1:] == to[i+1:]
		}
	}
	return false
}

func main() {

}
