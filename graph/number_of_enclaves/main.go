package main

import "fmt"

func dfs(grid [][]int, i int, j int, visited [][]bool, rows int, cols int) {
	visited[i][j] = true
	//visit up
	if i-1 >= 0 && grid[i-1][j] == 1 && !visited[i-1][j] {
		dfs(grid, i-1, j, visited, rows, cols)
	}
	//visit down
	if i+1 < rows && grid[i+1][j] == 1 && !visited[i+1][j] {
		dfs(grid, i+1, j, visited, rows, cols)
	}
	//visit left
	if j-1 >= 0 && grid[i][j-1] == 1 && !visited[i][j-1] {
		dfs(grid, i, j-1, visited, rows, cols)
	}
	//visit right
	if j+1 < cols && grid[i][j+1] == 1 && !visited[i][j+1] {
		dfs(grid, i, j+1, visited, rows, cols)
	}
}
func numberOfEnclaves(grid [][]int) int {
	//(0,0),(0,1),(0,2) first row
	//(0,0),(1,0),(2,0) first col
	//(2,0),(2,1),(2,2) last row
	//(0,2),(1,2),(2,2) last col
	rows := len(grid)
	cols := len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	for col := 0; col < cols; col++ {
		if grid[0][col] == 1 {
			dfs(grid, 0, col, visited, rows, cols)
		}
		if grid[rows-1][col] == 1 {
			dfs(grid, rows-1, col, visited, rows, cols)
		}
	}
	for row := 0; row < rows; row++ {
		if grid[row][0] == 1 {
			dfs(grid, row, 0, visited, rows, cols)
		}
		if grid[row][cols-1] == 1 {
			dfs(grid, row, cols-1, visited, rows, cols)
		}
	}
	enclaves := 0
	for i := range rows {
		for j := range cols {
			if grid[i][j] == 1 && !visited[i][j] {
				enclaves++
			}
		}
	}
	return enclaves

}
func main() {
	grid := [][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}
	fmt.Println(numberOfEnclaves(grid))
}
