package main

type Trie struct {
	isEnd bool
	node  map[byte]*Trie
}

func Constructor() Trie {
	node := make(map[byte]*Trie)
	trie := Trie{node: node, isEnd: false}
	return trie
}

func (this *Trie) Insert(word string) {
	prev := this

	for _, b := range word {
		trie := prev.node[byte(b)]
		if trie == nil {
			c := Constructor()
			// fmt.Println(c, byte(b))
			trie = &c
			prev.node[byte(b)] = trie
		}
		// fmt.Println(prev)
		prev = trie
	}
	prev.isEnd = true
	// fmt.Println(prev)
}

func dfs(board [][]byte, visited [][]bool, str string, x, y int, trie *Trie, res *[]string) {
	curByte := board[x][y]
	// fmt.Println(curByte)
	if visited[x][y] || trie.node[curByte] == nil {
		return
	}

	if trie.node[curByte] != nil {
		visited[x][y] = true
		str += string(curByte)
		// fmt.Println("str is ", str, " is End is ", trie, trie.node[curByte].isEnd)
		if trie.node[curByte].isEnd {
			*res = append(*res, str)
			trie.node[curByte].isEnd = false
			// return
		}
	}

	if x > 0 {
		dfs(board, visited, str, x-1, y, trie.node[curByte], res)
	}
	if x < len(board)-1 {
		dfs(board, visited, str, x+1, y, trie.node[curByte], res)
	}

	if y > 0 {
		dfs(board, visited, str, x, y-1, trie.node[curByte], res)
	}

	if y < len(board[0])-1 {
		dfs(board, visited, str, x, y+1, trie.node[curByte], res)
	}
	visited[x][y] = false
}

func findWords(board [][]byte, words []string) []string {
	trie := Constructor()
	for _, word := range words {
		trie.Insert(word)
	}

	boardSize, rowSize := len(board), len(board[0])

	visited := make([][]bool, boardSize, boardSize)
	for i := range visited {
		visited[i] = make([]bool, rowSize, rowSize)
	}

	// fmt.Println("visited is ", visited)
	var result []string
	for x := 0; x < boardSize; x++ {
		for y := 0; y < rowSize; y++ {
			dfs(board, visited, "", x, y, &trie, &result)
		}
	}

	return result
}

// func main() {

// 	// board := [][]byte{
// 	// 	{'o', 'a', 'a', 'n'},
// 	// 	{'e', 't', 'a', 'e'},
// 	// 	{'i', 'h', 'k', 'r'},
// 	// 	{'i', 'f', 'l', 'v'}}

// 	// words := []string{"oath", "pea", "eat", "rain"}

// 	board := [][]byte{{'a', 'a'}}
// 	words := []string{"a"}
// 	fmt.Println(findWords(board, words))
// }
