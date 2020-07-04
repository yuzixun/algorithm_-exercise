package main

import "fmt"

func wallsAndGates(rooms [][]int) {
	MAX := 1<<31 - 1
	fmt.Println(MAX)
	rows := len(rooms)
	if rows == 0 {
		return
	}
	columns := len(rooms[0])
	directions := [][]int{
		[]int{0, 1},
		[]int{0, -1},
		[]int{1, 0},
		[]int{-1, 0},
	}

	queue := [][2]int{}

	for i := 0; i < len(rooms); i++ {
		for j := 0; j < len(rooms[0]); j++ {
			if rooms[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	// fmt.Println(queue)
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]

		val := rooms[cur[0]][cur[1]]
		for _, direction := range directions {

			row := cur[0] + direction[0]
			column := cur[1] + direction[1]

			if row < 0 || row >= rows || column < 0 || column >= columns || MAX != rooms[row][column] {
				continue
			}
			rooms[row][column] = min(rooms[row][column], val+1)
			queue = append(queue, [2]int{row, column})
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	rooms := [][]int{
		[]int{2147483647, -1, 0, 2147483647},
		[]int{2147483647, 2147483647, 2147483647, -1},
		[]int{2147483647, -1, 2147483647, -1},
		[]int{0, -1, 2147483647, 2147483647},
	}
	wallsAndGates(rooms)
	fmt.Println(rooms)
}
