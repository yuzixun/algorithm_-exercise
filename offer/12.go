package main

func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	rows, columns := len(board), len(board[0])
	visited := make(map[int]bool, rows*columns)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if dfs(board, word, visited, rows, columns, i, j, 0) {
				return true
			}
			// fmt.Println(visited)
		}
	}
	return false
}

func dfs(board [][]byte, word string, visited map[int]bool, rows, columns, row, column, path int) bool {
	if path == len(word) {
		return true
	}

	hasPath := false
	x := columns*row + column
	if !visited[x] && row >= 0 && column >= 0 && row < rows && column < columns && board[row][column] == word[path] {
		path++
		visited[x] = true
		hasPath = dfs(board, word, visited, rows, columns, row-1, column, path) ||
			dfs(board, word, visited, rows, columns, row+1, column, path) ||
			dfs(board, word, visited, rows, columns, row, column-1, path) ||
			dfs(board, word, visited, rows, columns, row, column+1, path)

		if !hasPath {
			path--
			visited[x] = false
		}
	}

	return hasPath
}

// func main() {
// 	fmt.Println(exist([][]byte{}, "ABCCED"))
// 	fmt.Println(exist([][]byte{[]byte{'A', 'B', 'C', 'E'}, []byte{'S', 'F', 'C', 'S'}, []byte{'A', 'D', 'E', 'E'}}, "ABCCED"))
// 	fmt.Println(exist([][]byte{[]byte{'a', 'b'}, []byte{'c', 'd'}}, "abcd"))
// 	fmt.Println(exist([][]byte{[]byte{'C', 'A', 'A'}, []byte{'A', 'A', 'A'}, []byte{'B', 'C', 'D'}}, "AAB"))
// 	fmt.Println(exist([][]byte{[]byte{'A', 'B', 'C', 'E'}, []byte{'S', 'F', 'E', 'S'}, []byte{'A', 'D', 'E', 'E'}}, "ABCESEEEFS"))

// }
