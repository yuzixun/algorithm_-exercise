package main

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	rows := len(image)
	if rows == 0 || image[sr][sc] == newColor {
		return image
	}
	columns := len(image[0])
	dfs(&image, sr, sc, rows, columns, image[sr][sc], newColor)
	return image
}

func dfs(image *[][]int, i, j, rows, columns, curColor, newColor int) {
	if curColor == newColor {
		return
	}
	if i >= 0 && i <= rows-1 && j >= 0 && j <= columns-1 && curColor == (*image)[i][j] {
		(*image)[i][j] = newColor
		dfs(image, i-1, j, rows, columns, curColor, newColor)
		dfs(image, i+1, j, rows, columns, curColor, newColor)
		dfs(image, i, j-1, rows, columns, curColor, newColor)
		dfs(image, i, j+1, rows, columns, curColor, newColor)
	}

}

func main() {
	floodFill([][]int{[]int{0, 0, 0}, []int{0, 1, 1}}, 1, 1, 1)
}
