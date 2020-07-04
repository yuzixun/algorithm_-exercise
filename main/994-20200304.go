package main

func orangesRotting(grid [][]int) int {
	dx, dy := []int{1, -1, 0, 0}, []int{0, 0, 1, -1}
	size := len(grid)
	if size == 0 {
		return 0
	}

	result := 0
	rots := [][]int{}
	// 找到2的点
	for i := 0; i < size; i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				rots = append(rots, []int{i, j})
			}
		}
	}

	// BFS
	for len(rots) != 0 {
		nextRots := [][]int{}
		for rotIndex := 0; rotIndex < len(rots); rotIndex++ {
			rot := rots[rotIndex]
			for i := 0; i < 4; i++ {
				cx := rot[0] + dx[i]
				cy := rot[1] + dy[i]
				if cx >= 0 && cx < size && cy >= 0 && cy < len(grid[0]) && grid[cx][cy] == 1 {
					nextRots = append(nextRots, []int{cx, cy})
					grid[cx][cy] = -2
				}
			}
		}
		if len(nextRots) == 0 {
			break
		}
		// fmt.Println("rots", rots, nextRots)
		rots = nextRots
		result++
	}
	// fmt.Println(grid)
	for i := 0; i < size; i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	// fmt.Println("result is ", result)
	return result
}

func main() {
	// fmt.Println(orangesRotting([][]int{[]int{2, 1, 1}, []int{1, 1, 0}, []int{0, 1, 1}}))
	// fmt.Println(orangesRotting([][]int{[]int{2, 1, 1}, []int{0, 1, 1}, []int{1, 0, 1}}))
	// fmt.Println(orangesRotting([][]int{[]int{0, 2}}))
}
